package main

import (
	"log"
	"net/http"
)

func main() {

	trie := CreateNameTrie()

	fs := http.FileServer(http.Dir("frontend"))

	http.Handle("/", fs)

	http.HandleFunc("/api/card", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		log.Println(trie.getFuzzy(name))
	})
	log.Println("Starting server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
