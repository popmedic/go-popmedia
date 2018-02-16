package handle

import (
	"fmt"
	"net/http"
	"path"

	"github.com/popmedic/go-popmedia/server/context"
	"github.com/popmedic/go-popmedia/server/tmpl"
)

type Handshake struct {
	path    string
	context *context.Context
}

func NewHandshake(ctx *context.Context) *Handshake {
	return &Handshake{
		context: ctx,
	}
}

func (h *Handshake) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	return len(r.URL.Query()["handshake"]) > 0
}

func (h *Handshake) Handle(w http.ResponseWriter, r *http.Request) {
	tmpl, err := tmpl.LoadTemplate("Handshake", "templates/handshake.html")
	if nil != err {
		fmt.Fprint(w, "<<handshake>shakey shakey</handshake>")
		return
	}
	tmpl.Execute(w, nil)
}
