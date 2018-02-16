package handle

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/popmedic/go-popmedia/server/context"
	"github.com/popmedic/go-popmedia/server/info"
	"github.com/popmedic/go-popmedia/server/tmpl"
	"github.com/popmedic/go-wout/wout"
)

type Main struct {
	path    string
	context *context.Context
}

func NewMain(ctx *context.Context) *Main {
	return &Main{
		context: ctx,
	}
}

func (h *Main) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	fi, err := os.Stat(h.context.Config.Root + h.path)
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

	v, err := info.NewFilesAndDirectoriesInfoFromPath(h.context, h.path)
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
