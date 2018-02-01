package main

import (
	"flag"
	"log"

	"github.com/popmedic/popmedia2/server"
	"github.com/popmedic/popmedia2/server/config"
)

func main() {
	configPtr := flag.String("config", "./config.json", "path to config.json file")
	flag.Parse()

	err := config.MainConfig.LoadConfig(*configPtr)
	if nil != err {
		log.Println(err)
	}

	if err := server.Run(); nil != err {
		log.Println(err)
	}
}
