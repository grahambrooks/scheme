package main

import (
	"flag"
	"github.com/grahambrooks/scheme/service/server"
	"github.com/grahambrooks/scheme/service/store"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "sets the server port value")
	flag.Parse()

	scheme := server.NewSchemeServer(port, store.NewApiStore())

	scheme.ListenAndServe()
}
