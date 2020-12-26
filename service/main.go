package main

import "flag"

func main() {
	port := 8000
	flag.IntVar(&port, "port", 8000, "sets the server port value")

	flag.Parse()

	server := SchemeServer{Port: port, ApiStore: NewApiStore()}

	server.ListenAndServe()
}
