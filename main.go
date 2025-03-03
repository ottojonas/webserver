package main 

import ("fmt", "log", "net/http") 

// * hello handler with a simple message 
func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "wagwan")
}

func main(){
	// * set up the route for "/"
	http.HandleFunc("/", helloHandler)

	// * define the port to listen on 
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)

	// * start the server and log any errors
	log.Fatal(http.ListenAndServer(port, nill))
}


