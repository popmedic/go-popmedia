package info

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/popmedic/popmedia2/server/context"
	// "github.com/popmedic/popmedia2/server/config"
	// "github.com/popmedic/popmedia2/server/search"
)

type Mp4Info struct {
	Name        string
	Artist      string
	ReleaseDate string
	Genre       string
}

func NewMp4Info(ctx *context.Context, info *Info) (*Mp4Info, error) {
	cmd := exec.Command("/usr/local/bin/mp4info", filepath.Join(ctx.Config.Root, info.Path))
	res, err := cmd.Output()
	if nil != err {
		return nil, errors.New("unable to read mp4info cli output " + err.Error())
	}
	mp4Info := Mp4Info{}
	lines := strings.Split(string(res), "\n")
	for _, line := range lines {
		kv := strings.Split(line, ":")
		if len(kv) == 2 {
			k := strings.TrimSpace(kv[0])
			v := strings.TrimSpace(kv[1])
			switch k {
			case "Name":
				mp4Info.Name = v
			case "Artist":
				mp4Info.Artist = strings.Join(strings.Split(v, "|"), ", ")
			case "Release Date":
				mp4Info.ReleaseDate = v
			case "Genre":
				mp4Info.Genre = strings.Join(strings.Split(v, "|"), ", ")
			}
		}
	}
	return &mp4Info, nil
}

type Info struct {
	Name    string
	Path    string
	Host    string
	Root    string
	Image   string
	Desc    string
	Bif     string
	ExtInfo *Mp4Info
	context *context.Context
}

func (info *Info) LoadExtInfo() error {
	mp4Info, err := NewMp4Info(info.context, info)
	if nil != err {
		return err
	}
	info.ExtInfo = mp4Info
	return nil
}

type InfoList []*Info

func (il InfoList) Len() int {
	return len(il)
}

func (il InfoList) Less(i, j int) bool {
	return strings.TrimPrefix(strings.ToLower(il[i].Name), "the ") < strings.TrimPrefix(strings.ToLower(il[j].Name), "the ")
}

// Swap swaps the elements with indexes i and j.
func (il InfoList) Swap(i, j int) {
	il[i], il[j] = il[j], il[i]
}

func NewInfo(ctx *context.Context, name, path string) *Info {
	return &Info{
		Name:    name,
		Path:    path,
		Host:    host(ctx),
		Bif:     bif(path),
		Image:   image(ctx, path),
		Desc:    desc(ctx, path),
		context: ctx,
	}
}

type FilesAndDirectoriesInfo struct {
	Info        *Info
	Files       InfoList
	Directories InfoList
}

func NewFilesAndDirectoriesInfoFromPath(ctx *context.Context, path string) (*FilesAndDirectoriesInfo, error) {
	v := FilesAndDirectoriesInfo{
		Info:        NewInfo(ctx, strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)), path),
		Files:       InfoList{},
		Directories: InfoList{},
	}

	infos, err := getFileInfos(ctx, path)
	if nil != err {
		return nil, err
	}
	if path == "/" || path == "" {
		v.Info.Name = "PoPMediA"
	}
	for _, inf := range infos {
		if !strings.HasPrefix(inf.Name(), ".") &&
			!strings.HasPrefix(inf.Name(), "_") {
			inf = followSymLink(filepath.Join(ctx.Config.Root, path), inf)

			if inf.IsDir() {
				i := NewInfo(ctx, strings.TrimSuffix(inf.Name(), filepath.Ext(inf.Name())),
					filepath.Join(path, inf.Name()))
				v.Directories = append(v.Directories, i)
			} else if stringsContain(ctx.Config.MediaExt, strings.ToLower(filepath.Ext(inf.Name()))) {
				i := NewInfo(ctx, strings.TrimSuffix(inf.Name(), filepath.Ext(inf.Name())),
					filepath.Join(path, inf.Name()))
				v.Files = append(v.Files, i)
			}
		}
	}
	v.sortAll()
	return &v, nil
}

func NewFilesAndDirectoriesInfoFromSearch(ctx *context.Context, q string) *FilesAndDirectoriesInfo {
	v := FilesAndDirectoriesInfo{
		Info:        NewInfo(ctx, q, ""),
		Files:       InfoList{},
		Directories: InfoList{},
	}

	res := ctx.Search.Query(q)

	for path, name := range res {
		inf, err := os.Lstat(path)
		if err == nil {
			inf = followSymLink(path, inf)
			if inf.IsDir() {
				v.Directories = append(v.Directories, NewInfo(ctx, name, strings.TrimPrefix(path, ctx.Config.Root)+"/"))
			} else if stringsContain(ctx.Config.MediaExt, strings.ToLower(filepath.Ext(inf.Name()))) {
				v.Files = append(v.Files, NewInfo(ctx, name, strings.TrimPrefix(path, ctx.Config.Root)))
			}
		} else {
			log.Println(err)
		}
	}
	v.sortAll()
	return &v
}

func getFileInfos(ctx *context.Context, path string) ([]os.FileInfo, error) {
	f, err := os.Open(filepath.Join(ctx.Config.Root, path))
	if nil != err {
		return nil, err
	}
	defer func(f *os.File) {
		if err := f.Close(); nil != err {
			log.Println("Unable to close file: " + f.Name() + "\nError: " + err.Error())
		}
	}(f)

	return f.Readdir(0)
}

func (v *FilesAndDirectoriesInfo) sortAll() {
	sort.Sort(v.Files)
	sort.Sort(v.Directories)
}

func followSymLink(p string, info os.FileInfo) os.FileInfo {
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		_p, err := filepath.EvalSymlinks(filepath.Join(p, info.Name()))
		if nil == err {
			_info, err := os.Lstat(_p)
			if nil == err {
				return _info
			}
		}
	}
	return info
}

func host(ctx *context.Context) string {
	port := ctx.Config.Port
	if len(port) == 0 {
		return ctx.Config.Host
	}
	return ctx.Config.Host + ":" + port
}

func image(ctx *context.Context, path string) string {
	i := strings.TrimSuffix(path, filepath.Ext(path)) + "-SD.jpg"
	if _, err := os.Stat(filepath.Join(ctx.Config.Root, i)); err != nil {
		i = strings.TrimSuffix(i, filepath.Ext(i)) + ".png"
		if _, err := os.Stat(filepath.Join(ctx.Config.Root, i)); err != nil {
			var isDir = false
			if inf, err := os.Stat(filepath.Join(ctx.Config.Root, path)); nil == err {
				isDir = inf.IsDir()
			}

			if isDir {
				i = ctx.Config.DirectoryImage
			} else {
				i = ctx.Config.FileImage
			}
		}
	}
	return i
}

func bif(path string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + "-SD.bif"
}

func desc(ctx *context.Context, path string) string {
	p := filepath.Join(ctx.Config.Root, strings.TrimSuffix(path, filepath.Ext(path))+".desc")
	if _, err := os.Stat(p); err == nil {
		d, err := ioutil.ReadFile(p)
		if nil == err {
			return string(d)
		}
	}
	return ""
}

func stringsContain(ss []string, s string) bool {
	for _, str := range ss {
		if strings.ToLower(s) == strings.ToLower(str) {
			return true
		}
	}
	return false
}
