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

type Images struct {
	path    string
	context *context.Context
	re      *regexp.Regexp
}

func NewImages(ctx *context.Context) *Images {
	return &Images{
		context: ctx,
		re:      regexp.MustCompile(`^/images/.*\.[jJpP][pPnN][gG]$`),
	}
}

func (h *Images) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	return h.re.MatchString(strings.ToLower(h.path))
}

func (h *Images) Handle(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(h.path[1:])
	if nil != err {
		wout.Wout{err}.Print(w, "Unable to open image")
		return
	}

	defer file.Close()

	http.ServeContent(w, r, h.path, time.Now(), file)
}
