package context

import (
	"sync"

	"github.com/popmedic/popmedia2/server/config"
	"github.com/popmedic/popmedia2/server/search"
)

type Context struct {
	Config *config.Config
	Search *search.Search
	lock   sync.RWMutex
}

func NewContext() *Context {
	return &Context{
		Config: nil,
		Search: nil,
		lock:   sync.RWMutex{},
	}
}

func (ctx *Context) WithConfig(cfg *config.Config) *Context {
	ctx.SetConfig(cfg)
	return ctx
}

func (ctx *Context) SetConfig(cfg *config.Config) {
	ctx.lock.RLock()
	defer ctx.lock.RUnlock()
	ctx.Config = cfg
}

func (ctx *Context) WithSearch(s *search.Search) *Context {
	ctx.SetSearch(s)
	return ctx
}

func (ctx *Context) SetSearch(s *search.Search) {
	ctx.lock.RLock()
	defer ctx.lock.RUnlock()
	ctx.Search = s
}
