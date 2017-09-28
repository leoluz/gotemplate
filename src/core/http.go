package core

import (
	"io"
	"log"
	"net/http"
)

func StartServer() {
	log.Println("Starting server on port 8080...")
	http.HandleFunc("/", EchoHandler)
	http.ListenAndServe(":8080", nil)
}

func EchoHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello world!\n")
}
