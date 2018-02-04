package mux

import (
	"net/http"
	"sync"
)

type Muxer struct {
	defaultHandler      IHandler
	defaultHandlerMutex sync.RWMutex
	handlers            []IHandler
	handlersMutex       sync.RWMutex
}

func NewMuxer() *Muxer {
	return &Muxer{
		defaultHandler:      &notImplemented{},
		defaultHandlerMutex: sync.RWMutex{},
		handlers:            []IHandler{},
		handlersMutex:       sync.RWMutex{},
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

func (m *Muxer) ListenAndServe(addr string,
	listenAndServe func(addr string, handler Handler) error) error {
	return listenAndServe(addr, http.HandlerFunc(m.Handle))
}

type notImplemented struct{}

func (ni *notImplemented) Is(r *http.Request) bool {
	return true
}
func (ni *notImplemented) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("NOT IMPLEMENTED, default default handler."))
}
