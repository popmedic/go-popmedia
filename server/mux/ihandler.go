package mux

import (
	"net/http"
)

type IHandler interface {
	Is(*http.Request) bool
	Handle(http.ResponseWriter, *http.Request)
}
