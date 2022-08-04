package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	fmt.Println("Booting up on :3000")
	http.HandleFunc("/", helloHandler) // Update this line of code
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
