package server

import (
	"log"

	"github.com/popmedic/popmedia2/server/search"

	"github.com/popmedic/popmedia2/server/config"
	"github.com/popmedic/popmedia2/server/context"
	"github.com/popmedic/popmedia2/server/handle"
	"github.com/popmedic/popmedia2/server/mux"
)

func Run(ctx *context.Context) error {
	log.Println("Serving on port", config.MainConfig.Port, "with root", config.MainConfig.Root)

	ctx.SetSearch(search.MainSearch())

	handlers := []mux.IHandler{
		handle.NewFavicon(ctx),
		handle.NewHandshake(ctx),
		handle.NewRoku(ctx),
		handle.NewSearch(ctx),
		handle.NewPlayer(ctx),
		handle.NewMp4(ctx),
		handle.NewImages(ctx),
		handle.NewH404(ctx),
		handle.NewMain(ctx),
	}

	muxer := mux.NewMuxer(ctx).WithHandlers(handlers).WithDefaultHandler(handle.NewDefault(ctx))
	return muxer.ListenAndServe(":" + config.MainConfig.Port)
}
