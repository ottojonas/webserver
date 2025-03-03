package main

import (
    "fmt"
    "log"
    "net/http"
)

// helloHandler responds with a simple message.
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "wagwan")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users endpoint")
}

func main() {
    // Set up the route for "/"
    http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api/users", usersHandler)

    // Define the port to listen on
    port := ":3001"
    fmt.Printf("Server is running on http://localhost%s\n", port)

    // Start the server and log any errors.
    log.Fatal(http.ListenAndServe(port, nil))
}