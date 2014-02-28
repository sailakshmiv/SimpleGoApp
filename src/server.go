package main

import (
	"os"
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "text/html")
    fmt.Fprintf(w, "<h2>Hello World from Distelli!</h2>")
}

func main() {
	port := os.Args[1]
	fmt.Println(port)

    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+port, nil)
}
