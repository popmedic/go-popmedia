package handle

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/popmedic/popmedia2/server/config"
	"github.com/popmedic/popmedia2/server/info"
	"github.com/popmedic/popmedia2/server/tmpl"
	"github.com/popmedic/wout"
)

type Main struct {
	path string
}

func NewMain() *Main {
	return &Main{}
}

func (h *Main) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	fi, err := os.Stat(config.MainConfig.Root + h.path)
	if nil != err {
		fmt.Println(err)
		return false
	}
	return fi.IsDir()
}

func (h *Main) Handle(w http.ResponseWriter, r *http.Request) {
	tmpl, err := tmpl.LoadTemplate("main-directory", "templates/main.html")
	if nil != err {
		wout.Wout{err}.Print(w, "unable to load template \"main-directory\"")
		return
	}

	v, err := info.NewFilesAndDirectoriesInfoFromPath(h.path)
	if nil != err {
		wout.Wout{err}.Print(w, "unable to get infos for path \""+h.path+"\"")
		return
	}

	err = tmpl.Execute(w, v)
	if nil != err {
		wout.Wout{err}.Print(w, "Unable to execute template in main")
		return
	}
}
