package handle

import (
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/popmedic/wout"
)

var imgRE = regexp.MustCompile(`^/images/.*\.[jJpP][pPnN][gG]$`)

type Images struct {
	path string
}

func NewImages() *Images {
	return &Images{}
}

func (h *Images) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	return imgRE.MatchString(strings.ToLower(h.path))
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
