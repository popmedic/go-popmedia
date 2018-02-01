package handle

import (
	"fmt"
	"net/http"
	"path"
	"regexp"
	"strings"

	"github.com/popmedic/popmedia2/server/tmpl"
	"github.com/popmedic/wout"

	"github.com/popmedic/popmedia2/server/info"
)

type Search struct {
	path string
}

func NewSearch() *Search {
	return &Search{}
}

func (h *Search) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	return regexp.MustCompile("^/search").MatchString(h.path)
}

func (h *Search) Handle(w http.ResponseWriter, r *http.Request) {
	tmpl, err := tmpl.LoadTemplate("main-search", "templates/main.html")
	if nil != err {
		fmt.Fprintf(w, err.Error())
		return
	}

	r.ParseForm()
	q := strings.Join(r.Form["q"], " ")

	v := info.NewFilesAndDirectoriesInfoFromSearch(q)

	err = tmpl.Execute(w, v)
	if nil != err {
		wout.Wout{err}.Print(w, "Unable to execute template in search")
		return
	}
}
