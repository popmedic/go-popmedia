package handle

import (
	"net/http"
	"path"

	"github.com/popmedic/popmedia2/server/context"
)

type Default struct {
	path    string
	context *context.Context
}

func NewDefault(ctx *context.Context) *Default {
	return &Default{
		context: ctx,
	}
}

func (h *Default) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	return true
}

func (h *Default) Handle(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir(h.context.Config.Root)).ServeHTTP(w, r)
}
