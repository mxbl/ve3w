package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	certFile := "cert/ve3w.crt"
	keyFile := "cert/ve3w.key"

	fmt.Println("Starting HTTPS server on port 8080")
	err := http.ListenAndServeTLS(":8080", certFile, keyFile, nil)
	if err != nil {
		log.Fatal("Error starting HTTPS server: ", err)
	}
}
