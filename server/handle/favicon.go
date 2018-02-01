package handle

import (
	"net/http"
	"os"
	"path"
	"regexp"
	"time"

	"github.com/popmedic/wout"
)

type Favicon struct {
	path string
}

func NewFavicon() *Favicon {
	return &Favicon{}
}

func (h *Favicon) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	return regexp.MustCompile("^/favicon.ico").MatchString(h.path)
}

func (h *Favicon) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Connection", "keep-alive")
	file, err := os.Open("templates" + h.path)
	if nil != err {
		wout.Wout{err}.Print(w, "Unable to open favicon")
		return
	}

	defer file.Close()

	http.ServeContent(w, r, h.path, time.Now(), file)
}
