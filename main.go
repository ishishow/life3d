package main

import (
	"flag"
	"lifegame/presenter/http/server"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {

	//err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	//if err != nil {
	//	log.Fatalf("getenv is failed! :%v", err)
	//}

	server.Serve(addr)
}
