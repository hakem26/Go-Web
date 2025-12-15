package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	CreaedAt time.Time `json: "created_at"`
}

type getUser interface {
	GetUsers() []User
}

func getUsers() []User {
	return []User {
		{Name: "Alice", Age: 30, CreaedAt: time.Now()},
		{Name: "Bob", Age: 25, CreaedAt: time.Now()},
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Hello World!</h1>"))
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.ParseForm() != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Form submitted successfully!\n")
	name := r.FormValue("name")
	age := r.FormValue("age")
	fmt.Fprintf(w, "Hello %s you are %s years old", name, age)
}

func main()  {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	
	fmt.Println("The server runs on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}