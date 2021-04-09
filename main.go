package main

import (
	"flag"
	"fmt"
	"lifegame/presenter/http/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	addr string
)

func init() {
	//flag.StringVar(&addr, "addr", ":"+os.Getenv("PORT"), "tcp host:port to connect")
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {

	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatalf("getenv is failed! :%v", err)
	}

	server.Serve(addr)
}
