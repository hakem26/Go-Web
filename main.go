package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte("<h1>Welcome to the Home Page</h1>"))
	
}

func main()  {
	mux := http.ListenAndServe(":8080", http.HandlerFunc(home))
	if err := mux; err != nil {
		log.Fatal(err)
	}
}