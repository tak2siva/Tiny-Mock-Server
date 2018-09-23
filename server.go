package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func responseCallback(w http.ResponseWriter, r *http.Request) {
	content := os.Getenv("CONTENT")
	if len(content) == 0 {
		content = "Hello from tiny mock server"
	}

	code := os.Getenv("CODE")
	if len(code) == 0 {
		code = "200"
	}

	i, _ := strconv.Atoi(code)
	w.WriteHeader(i)
	fmt.Fprintf(w, content) // send data to client side
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9090"
	}

	fmt.Printf("Started on port: %s\n", port)
	http.HandleFunc("/", responseCallback)                    // set router
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
