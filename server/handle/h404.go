package handle

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/popmedic/popmedia2/server/config"
	"github.com/popmedic/popmedia2/server/tmpl"
)

type H404 struct {
	path string
}

func NewH404() *H404 {
	return &H404{}
}

func (h *H404) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	fs := http.Dir(config.MainConfig.Root)
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
