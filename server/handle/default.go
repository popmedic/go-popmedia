package handle

import (
	"net/http"
	"path"

	"github.com/popmedic/popmedia2/server/config"
)

type Default struct {
	path string
}

func NewDefault() *Default {
	return &Default{}
}

func (h *Default) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	return true
}

func (h *Default) Handle(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir(config.MainConfig.Root)).ServeHTTP(w, r)
}
