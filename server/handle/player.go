package handle

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/popmedic/popmedia2/server/context"
	"github.com/popmedic/popmedia2/server/info"
	"github.com/popmedic/popmedia2/server/tmpl"
	"github.com/popmedic/wout"
)

type Player struct {
	path    string
	context *context.Context
	re      *regexp.Regexp
}

func NewPlayer(ctx *context.Context) *Player {
	return &Player{
		context: ctx,
		re:      regexp.MustCompile("^/player"),
	}
}

func (h *Player) Is(r *http.Request) bool {
	h.path = path.Clean(r.URL.Path)
	if h.re.MatchString(h.path) {
		h.path = strings.TrimLeft(h.path, "/player")
		return true
	}
	return false
}

func (h *Player) Handle(w http.ResponseWriter, r *http.Request) {
	tmpl, err := tmpl.LoadTemplate("player", "templates/player.html")
	if nil != err {
		fmt.Fprintf(w, err.Error())
		return
	}

	inf := info.NewInfo(h.context, strings.TrimSuffix(filepath.Base(h.path), filepath.Ext(h.path)), h.path)
	err = inf.LoadExtInfo()
	if nil != err {
		log.Println(err)
	}

	err = tmpl.Execute(w, inf)
	if nil != err {
		wout.Wout{err}.Print(w, "Unable to execute template in main")
		return
	}
}
