package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "🦄🦄🦄 FOO Awesome Hello World from Bar (Go)! ")
	fmt.Fprintf(w, hostname)
	fmt.Fprintf(w, "🦄🦄🦄\n")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server started.")
	http.ListenAndServe(":80", nil)
}
