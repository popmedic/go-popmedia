package handle

import (
	"net/http"
	"path"
	"strings"

	"github.com/popmedic/popmedia2/server/info"
	"github.com/popmedic/popmedia2/server/tmpl"
	"github.com/popmedic/wout"
)

type Roku struct {
	path string
}

func NewRoku() *Roku {
	return &Roku{}
}

func (h *Roku) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	if strings.ToLower(path.Base(h.path)) == "roku.php" {
		h.path = path.Dir(h.path)
		return true
	}
	return false
}

func (h *Roku) Handle(w http.ResponseWriter, r *http.Request) {
	tmplName := "roku-home"
	tmplPath := "templates/roku_home.xml"
	if h.isRokuDir(r) {
		tmplName = "roku-dir"
		tmplPath = "templates/roku_dir.xml"

		h.path = frontSlash(r.URL.Query()["dir"][0])
	}

	tmpl, err := tmpl.LoadTemplate(tmplName, tmplPath)
	if nil != err {
		wout.Wout{err}.Print(w, "unable to load template \""+tmplName+"\"")
		return
	}

	v, err := info.NewFilesAndDirectoriesInfoFromPath(h.path)
	if nil != err {
		wout.Wout{err}.Print(w, "unable to get infos for path \""+h.path+"\"")
		return
	}

	err = tmpl.Execute(w, v)
	if nil != err {
		wout.Wout{err}.Print(w, "Unable to execute template in \""+tmplName+"\"")
		return
	}
}

func (h *Roku) isRokuDir(r *http.Request) bool {
	return len(r.URL.Query()["dir"]) > 0 && r.URL.Query()["dir"][0] != "/"
}

func frontSlash(path string) string {
	if len(path) > 0 {
		if path[0] != '/' {
			return "/" + path
		}
	}
	return path
}
