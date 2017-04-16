package server

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

type Mp4Info struct {
	Name        string
	Artist      string
	ReleaseDate string
	Genre       string
}

func newMp4Info(info *Info) (*Mp4Info, error) {
	cmd := exec.Command("mp4info", filepath.Join(MainConfig.Root, info.Path))
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
	Image   string
	Desc    string
	ExtInfo *Mp4Info
}

func (info *Info) loadExtInfo() error {
	mp4Info, err := newMp4Info(info)
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

func newInfo(name, path string) *Info {
	info := Info{
		Name: name,
		Path: path,
	}
	info.Image = strings.TrimSuffix(path, filepath.Ext(path)) + "-SD.jpg"
	if _, err := os.Stat(filepath.Join(MainConfig.Root, info.Image)); err != nil {
		info.Image = MainConfig.DirectoryImage
	}

	p := filepath.Join(MainConfig.Root, strings.TrimSuffix(path, filepath.Ext(path))+".desc")
	if _, err := os.Stat(p); err == nil {
		d, err := ioutil.ReadFile(p)
		if nil == err {
			info.Desc = string(d)
		}
	}
	return &info
}

type FilesAndDirectoriesInfo struct {
	Info        *Info
	Files       InfoList
	Directories InfoList
}

func (v *FilesAndDirectoriesInfo) SortAll() {
	sort.Sort(v.Files)
	sort.Sort(v.Directories)
}
