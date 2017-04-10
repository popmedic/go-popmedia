package server

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func loadTemplate(name, path string) (*template.Template, error) {
	b, err := ioutil.ReadFile(path)
	if nil != err {
		return nil, err
	}
	funcMap := template.FuncMap{
		"replaceDashes": replaceDashes,
		"splitDashes":   splitDashes,
		"split":         strings.Split,
		"joinPath":      joinPath,
	}
	return template.New(name).Funcs(funcMap).Parse(string(b))
}
func respond404(p string, w http.ResponseWriter) {
	tmpl, err := loadTemplate("404", "templates/404.html")
	if nil != err {
		fmt.Fprintf(w, "404 you dumbass. you got and error:", err)
		return
	}
	tmpl.Execute(w, struct{ Path string }{Path: p})
	return
}

func respondMain(r, p string, w http.ResponseWriter) {
	tmpl, err := loadTemplate("main-directory", "templates/main.html")
	if nil != err {
		fmt.Fprintf(w, err.Error())
		return
	}

	f, err := os.Open(filepath.Join(r, p))
	if nil != err {
		fmt.Fprintf(w, err.Error())
		log.Println("Unable to execute open", err)
		return
	}

	infos, err := f.Readdir(0)
	if nil != err {
		fmt.Fprintf(w, err.Error())
		log.Println("Unable to execute readdir", err)
		return
	}
	v := FilesAndDirectoriesInfo{
		Info:        newInfo(strings.TrimSuffix(filepath.Base(p), filepath.Ext(p)), p),
		Files:       InfoList{},
		Directories: InfoList{},
	}
	if p == "/" || p == "" {
		v.Info.Name = "PoPMediA"
	}
	for _, info := range infos {
		if !strings.HasPrefix(info.Name(), ".") &&
			!strings.HasPrefix(info.Name(), "_") {
			info = followSymLink(filepath.Join(r, p), info)

			if info.IsDir() {
				i := newInfo(strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())),
					filepath.Join(p, info.Name()))
				v.Directories = append(v.Directories, i)
			} else if stringsContain(MainConfig.MediaExt, strings.ToLower(filepath.Ext(info.Name()))) {
				i := newInfo(strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())),
					filepath.Join(p, info.Name()))
				v.Files = append(v.Files, i)
			}
		}
	}
	v.SortAll()
	err = tmpl.Execute(w, v)
	if nil != err {
		fmt.Fprintf(w, err.Error())
		log.Println("Unable to execute template in main", err)
		return
	}
}

func respondSearch(root string, w http.ResponseWriter, r *http.Request) {
	tmpl, err := loadTemplate("main-search", "templates/main.html")
	if nil != err {
		fmt.Fprintf(w, err.Error())
		return
	}

	r.ParseForm()
	q := strings.Join(r.Form["q"], " ")
	v := FilesAndDirectoriesInfo{
		Info:        newInfo(q, ""),
		Files:       InfoList{},
		Directories: InfoList{},
	}

	res := MainSearch().Query(q)
	for path, name := range res {
		info, err := os.Lstat(path)
		if err == nil {
			info = followSymLink(path, info)
			if info.IsDir() {
				i := newInfo(name, strings.TrimPrefix(path, root))
				v.Directories = append(v.Directories, i)
			} else if stringsContain(MainConfig.MediaExt, strings.ToLower(filepath.Ext(info.Name()))) {
				i := newInfo(name, strings.TrimPrefix(path, root))
				v.Files = append(v.Files, i)
			}
		} else {
			log.Println(err)
		}
	}

	v.SortAll()
	err = tmpl.Execute(w, v)
	if nil != err {
		fmt.Fprintf(w, err.Error())
		log.Println("Unable to execute template in search", err)
		return
	}
}

func respondPlayer(p string, w http.ResponseWriter) {
	tmpl, err := loadTemplate("player", "templates/player.html")
	if nil != err {
		fmt.Fprintf(w, err.Error())
		return
	}

	info := newInfo(strings.TrimSuffix(filepath.Base(p), filepath.Ext(p)), p)

	err = tmpl.Execute(w, info)
	if nil != err {
		fmt.Fprintf(w, err.Error())
		log.Println("Unable to execute template in main", err)
		return
	}
}

func Run() error {
	cfg, err := newConfig("config.json")
	if nil != err {
		return err
	}
	log.Println("Serving on port", cfg.Port, "with root", cfg.Root)

	return http.ListenAndServe(":"+cfg.Port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs := http.Dir(cfg.Root)
		p := path.Clean(r.URL.Path)
		if isSearch(p) {
			respondSearch(cfg.Root, w, r)
		} else if isPlayer(p) {
			respondPlayer(strings.TrimPrefix(p, "/player"), w)
		} else {
			_, err := fs.Open(p)
			if os.IsNotExist(err) {
				respond404(p, w)
			} else {
				fi, err := os.Stat(cfg.Root + p)
				if nil != err {
					fmt.Fprintln(w, err)
				} else if fi.IsDir() {
					respondMain(cfg.Root, p, w)
				} else {
					http.FileServer(fs).ServeHTTP(w, r)
				}
			}
		}
	}))
}

func isSearch(p string) bool {
	return regexp.MustCompile("^/search").MatchString(p)
}

func isPlayer(p string) bool {
	return regexp.MustCompile("^/player").MatchString(p)
}

func isMp4(fi os.FileInfo) bool {
	return strings.ToLower(filepath.Ext(fi.Name())) == ".mp4"
}

func stringsContain(ss []string, s string) bool {
	for _, str := range ss {
		if s == str {
			return true
		}
	}
	return false
}

func replaceDashes(s string) template.HTML {
	return template.HTML(strings.Replace(s, " - ", "<br/>", -1))
}

func splitDashes(s string) []string {
	return strings.Split(s, " - ")
}

func joinPath(ss []string, n int) string {
	return string(filepath.Separator) + filepath.Join(ss[0:n+1]...)
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

func followSymLinkWithPath(p string, info os.FileInfo) (os.FileInfo, string) {
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		_p, err := filepath.EvalSymlinks(p)
		if nil == err {
			_info, err := os.Lstat(_p)
			if nil == err {
				return _info, _p
			}
		} else {
			log.Println(err)
		}
	}
	return info, p
}
