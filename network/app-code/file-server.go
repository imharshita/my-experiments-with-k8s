package main

import (
	"log"
	"net/http"
)

func main() {

	index := http.FileServer(http.Dir("/"))

	fileServer := http.FileServer(http.Dir("/serve/data/"))

	http.Handle("/", index)

	http.Handle("/files/", http.StripPrefix("/files", fileServer))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
