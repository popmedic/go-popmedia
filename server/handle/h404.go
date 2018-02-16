package handle

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/popmedic/go-popmedia/server/context"
	"github.com/popmedic/go-popmedia/server/tmpl"
)

type H404 struct {
	path    string
	context *context.Context
}

func NewH404(ctx *context.Context) *H404 {
	return &H404{
		context: ctx,
	}
}

func (h *H404) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	fs := http.Dir(h.context.Config.Root)
	_, err := fs.Open(h.path)
	return os.IsNotExist(err)
}

func (h *H404) Handle(w http.ResponseWriter, r *http.Request) {
	tmpl, err := tmpl.LoadTemplate("404", "templates/404.html")
	if nil != err {
		fmt.Fprintf(w, "404 you dumbass. you got and error: %q", err)
		return
	}
	tmpl.Execute(w, struct{ Path string }{Path: h.path})
	return
}
