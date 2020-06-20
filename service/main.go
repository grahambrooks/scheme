package main

func main() {
	server := SchemeServer{Port: 8000, ApiStore: NewApiStore()}

	server.ListenAndServe()
}
