package handle

import (
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/popmedic/go-popmedia/server/context"
	"github.com/popmedic/go-wout/wout"
)

type Mp4 struct {
	path    string
	context *context.Context
	re      *regexp.Regexp
}

func NewMp4(ctx *context.Context) *Mp4 {
	return &Mp4{
		context: ctx,
		re:      regexp.MustCompile(".mp4$"),
	}
}

func (h *Mp4) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	return h.re.MatchString(strings.ToLower(h.path))
}

func (h *Mp4) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Connection", "keep-alive")
	file, err := os.Open(h.context.Config.Root + h.path)
	if nil != err {
		wout.Wout{err}.Print(w, "Unable to open video")
		return
	}

	defer file.Close()
	modtime := time.Now()
	if fs, err := file.Stat(); nil != err {
		wout.Wout{err}.Print(w, "Unable to get stat on file "+h.path)
	} else {
		modtime = fs.ModTime()
	}
	http.ServeContent(w, r, h.path, modtime, file)
}
