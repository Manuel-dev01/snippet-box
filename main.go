package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()


	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("Server working!")
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	
	if err != nil && id < 1 {
		http.NotFound(w, r)
		return 
	}
	fmt.Println("Display a specific snippet")
	w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	fmt.Println("Create a new snippet")
	w.Write([]byte("Create a new snippet..."))
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	fmt.Println("display the home page")
	w.Write([]byte("Hello from Snippetbox"))
}