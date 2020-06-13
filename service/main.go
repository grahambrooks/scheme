package main

func main() {
	server := ApelliconServer{Port: 8000, ApiStore: NewApiStore()}

	server.ListenAndServe()
}
