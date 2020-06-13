package main

const (
	SearchIndexName   = "apellicon-search"
	DocumentIndexName = "apellicon-docs"
)

func main() {
	server := ApelliconServer{Port: 8000}

	server.ListenAndServe()
}
