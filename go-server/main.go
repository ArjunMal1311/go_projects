package main

import (
	"fmt"
	"log"
	"net/http"
)

type formHandler struct{}

func (h *formHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() error : %v", err)
		return
	}

	fmt.Fprintf(w, "POST request success")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s \n", name)
	fmt.Fprintf(w, "Address = %s \n", address)
}

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.Handle("/form", &formHandler{})
	http.Handle("/hello", &helloHandler{})

	fmt.Printf("Starting server at port 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
