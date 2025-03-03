package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "wagwan")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users endpoint")
}

func main() {
    http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api/users", usersHandler)
    port := ":3001"
    fmt.Printf("Server is running on http://localhost%s\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}