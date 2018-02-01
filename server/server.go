package server

import (
	"log"
	"net/http"

	"github.com/popmedic/popmedia2/server/config"
	"github.com/popmedic/popmedia2/server/handle"
	"github.com/popmedic/popmedia2/server/mux"
)

func Run() error {
	err := config.MainConfig.LoadConfig("config.json")
	if nil != err {
		return err
	}
	cfg := config.MainConfig
	log.Println("Serving on port", cfg.Port, "with root", cfg.Root)

	handlers := []mux.IHandler{
		handle.NewFavicon(),
		handle.NewHandshake(),
		handle.NewRoku(),
		handle.NewSearch(),
		handle.NewPlayer(),
		handle.NewMp4(),
		handle.NewH404(),
		handle.NewMain(),
	}

	muxer := mux.NewMuxer().WithHandlers(handlers).WithDefaultHandler(handle.NewDefault())

	return http.ListenAndServe(":"+cfg.Port, http.HandlerFunc(muxer.Handle))
}
