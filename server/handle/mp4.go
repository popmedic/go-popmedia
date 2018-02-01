package handle

import (
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/popmedic/popmedia2/server/config"
	"github.com/popmedic/wout"
)

type Mp4 struct {
	path string
}

func NewMp4() *Mp4 {
	return &Mp4{}
}

func (h *Mp4) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	return regexp.MustCompile(".mp4$").MatchString(strings.ToLower(h.path))
}

func (h *Mp4) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Connection", "keep-alive")
	file, err := os.Open(config.MainConfig.Root + h.path)
	if nil != err {
		wout.Wout{err}.Print(w, "Unable to open video")
		return
	}

	defer file.Close()

	http.ServeContent(w, r, h.path, time.Now(), file)
}
