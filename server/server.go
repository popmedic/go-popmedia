package server

import (
	"net/http"

	"github.com/popmedic/go-logger/log"
	"github.com/popmedic/go-mux/mux"
	"github.com/popmedic/go-popmedia/server/config"
	"github.com/popmedic/go-popmedia/server/context"
	"github.com/popmedic/go-popmedia/server/handle"
)

func Run(ctx *context.Context, listenAndServe func(addr string, handler http.Handler) error) error {
	log.Info("Serving on port", ctx.Config.Port, "with root", ctx.Config.Root)

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

	muxer := mux.NewMuxer().WithHandlers(handlers).WithDefaultHandler(handle.NewDefault(ctx))
	return muxer.ListenAndServe(":"+config.MainConfig.Port, listenAndServe)
}
