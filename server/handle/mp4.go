package handle

import (
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/popmedic/popmedia2/server/context"
	"github.com/popmedic/wout"
)

var mp4RE = regexp.MustCompile(".mp4$")

type Mp4 struct {
	path    string
	context *context.Context
}

func NewMp4(ctx *context.Context) *Mp4 {
	return &Mp4{
		context: ctx,
	}
}

func (h *Mp4) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	return mp4RE.MatchString(strings.ToLower(h.path))
}

func (h *Mp4) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Connection", "keep-alive")
	file, err := os.Open(h.context.Config.Root + h.path)
	if nil != err {
		wout.Wout{err}.Print(w, "Unable to open video")
		return
	}

	defer file.Close()

	http.ServeContent(w, r, h.path, time.Now(), file)
}
