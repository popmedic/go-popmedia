package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"github.com/popmedic/popmedia2/server/search"

	"github.com/popmedic/popmedia2/server"
	"github.com/popmedic/popmedia2/server/config"
	"github.com/popmedic/popmedia2/server/context"
)

func main() {
	configPtr := flag.String("config", "./config.json", "path to config.json file")
	flag.Parse()

	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		log.Fatal(err)
	} else {
		if err := os.Chdir(dir); err != nil {
			log.Fatal(err)
		}
	}

	if err := config.MainConfig.LoadConfig(*configPtr); nil != err {
		log.Println(err)
	}

	if len(config.MainConfig.LogFile) > 0 {
		if lout, err := os.OpenFile(config.MainConfig.LogFile,
			syscall.O_CREAT|syscall.O_APPEND|syscall.O_WRONLY,
			os.ModePerm); err != nil {
			log.Println(err)
		} else {
			log.SetOutput(lout)
		}

	}

	ctx := context.NewContext().WithConfig(config.MainConfig).WithSearch(search.MainSearch())

	if err := server.Run(ctx); nil != err {
		log.Println(err)
	}
}
