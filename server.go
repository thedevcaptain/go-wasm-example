package main

import "net/http"

func main() {
	// serve the static files into public folder
	http.Handle("/", http.FileServer(http.Dir("public")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
