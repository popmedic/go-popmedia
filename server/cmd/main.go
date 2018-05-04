package main

import (
	"flag"
	"net/http"
	"os"
	"path/filepath"
	"syscall"

	"github.com/popmedic/go-logger/log"
	"github.com/popmedic/go-popmedia/server"
	"github.com/popmedic/go-popmedia/server/config"
	"github.com/popmedic/go-popmedia/server/context"
	"github.com/popmedic/go-popmedia/server/search"
)

func main() {
	configPtr := flag.String("config", "./config.json", "path to config.json file")
	flag.Parse()

	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		log.Fatal(os.Exit, err)
	} else {
		if err := os.Chdir(dir); err != nil {
			log.Fatal(os.Exit, err)
		}
	}

	if err := config.MainConfig.LoadConfig(*configPtr); nil != err {
		log.Error(err)
	}

	if len(config.MainConfig.LogFile) > 0 {
		if lout, err := os.OpenFile(config.MainConfig.LogFile,
			syscall.O_CREAT|syscall.O_APPEND|syscall.O_WRONLY,
			os.ModePerm); err != nil {
			log.Error(os.Exit, err)
		} else {
			log.SetOutput(lout)
		}
	}

	if config.MainConfig.UseHTMLLogger {
		log.SetHTMLStatus(true, ":"+config.MainConfig.HTMLLoggerPort)
	}

	ctx := context.NewContext().WithConfig(config.MainConfig).WithSearch(search.MainSearch(config.MainConfig))

	if err := server.Run(ctx, http.ListenAndServe); nil != err {
		log.Info(err)
	}
}
