package server

import (
	"log"

	"github.com/popmedic/popmedia2/server/config"
	"github.com/popmedic/popmedia2/server/handle"
	"github.com/popmedic/popmedia2/server/mux"
)

func Run() error {
	log.Println("Serving on port", config.MainConfig.Port, "with root", config.MainConfig.Root)

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
	return muxer.ListenAndServe(":" + config.MainConfig.Port)
}
