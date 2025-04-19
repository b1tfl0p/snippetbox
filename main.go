package main

import (
	"log"
	"net/http"
)

/* Three ESSENTIALS for a web app:
* 1. Functions (Handlers) that write HTTP response headers & bodies
* 2. A router to map paths to function handlers
* 3. A web server to listen for incoming requests to hand to the router
 */

// 1. A "home" handler function.
// Responsible for executing your application logic and for writing HTTP
// response headers and bodies.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// 2. http.NewServeMux() function initializes a server multiplexer.
	// Use the servemux to register URL patterns with the function you want to
	// use to handle the requests to that path
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// 3. The web server
	// Use the http.ListenAndServe() function to start a new web server.
	// The parameters determine:
	// * the TCP network address to listen on
	// * the function (router) to use to handle requests
	port := ":4000"
	log.Printf("Starting server on %s", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
