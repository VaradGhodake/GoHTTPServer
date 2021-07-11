package main

// import necessary packages for logging and serving 
// HTTP requests
import (
	"log"
	"net/http"
	"time"
)

// wrap our http.Handler
func myMiddleware(handler http.Handler) http.Handler {
    // here, use anonymous function syntax and return an http.Handler object
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // before processing
            log.Println("Started req processing at ", time.Now().Format(time.RFC850))

            // let's sleep for 5 seconds to simulate middleware functionality
            time.Sleep(5 * time.Second)
            // http.handler will have this method to satisfy the Handler interface
            handler.ServeHTTP(w, r)

            // after processing
            log.Println("Ended req processing at ", time.Now().Format(time.RFC850))
    })
}

// function that handles /foo
func foo_function(w http.ResponseWriter, r *http.Request) {
	// business logic
    w.Write([]byte("called foo endpoint through middleware")) // write with typecasting into bytes
}


func main() {
    foo_function_handler := http.HandlerFunc(foo_function)
    
    // we won't use http.HandleFunc because the middleware is returning an http.Handler object
    http.Handle("/foo", myMiddleware(foo_function_handler))
    log.Fatal(http.ListenAndServe(":3000", nil))
}
