package main

// import necessary packages for logging and serving 
// HTTP requests
import (
	"log"
	"net/http"
)

func foo_function(w http.ResponseWriter, r *http.Request) {
	// business logic
    w.Write([]byte("called foo endpoint")) // write with typecasting into bytes
}

func main() {
    // foo_function_handler = http.HandlerFunc(foo_function)
    
    // we won't use http.HandleFunc because the middleware is giving us an http.Handler object
    http.HandleFunc("/foo", foo_function)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
