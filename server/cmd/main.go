package main

import (
	"github.com/popmedic/popmedia2/server"
	"log"
)

func main(){
	if err := server.Run(); nil != err {
		log.Println(err)
	}
}