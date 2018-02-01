package mux

import (
	"net/http"
	"sync"
)

type Muxer struct {
	handlers            []IHandler
	defaultHandler      IHandler
	handlersMutex       sync.RWMutex
	defaultHandlerMutex sync.RWMutex
}

func NewMuxer() *Muxer {
	return &Muxer{
		handlersMutex:       sync.RWMutex{},
		defaultHandlerMutex: sync.RWMutex{},
	}
}

func (m *Muxer) With(handlers []IHandler, defaultHandler IHandler) *Muxer {
	return m.WithHandlers(handlers).WithDefaultHandler(defaultHandler)
}

func (m *Muxer) WithHandlers(handlers []IHandler) *Muxer {
	m.handlersMutex.Lock()
	defer m.handlersMutex.Unlock()

	m.handlers = make([]IHandler, len(handlers))
	copy(m.handlers, handlers)
	return m
}

func (m *Muxer) WithDefaultHandler(defaultHandler IHandler) *Muxer {
	m.defaultHandlerMutex.Lock()
	defer m.defaultHandlerMutex.Unlock()

	m.defaultHandler = defaultHandler
	return m
}

func (m *Muxer) Handle(w http.ResponseWriter, r *http.Request) {
	m.handlersMutex.RLock()
	defer m.handlersMutex.RUnlock()

	for _, h := range m.handlers {
		if h.Is(r) {
			h.Handle(w, r)
			return
		}
	}

	m.defaultHandlerMutex.RLock()
	defer m.defaultHandlerMutex.RUnlock()

	m.defaultHandler.Handle(w, r)
}
