package main

import (
	"fmt"
	//import "fmt" short for  format
	"net/http"

	"github.com/gorilla/mux"

	// import handler
	"github.com/jazz/handlers"
)

func main() {
	fmt.Printf("HelloWorld!\n")
	// new router is to understand which handler to execute or call
	r := mux.NewRouter()

	// we need to import the handler
	// when running the browser you must type in localhost:3000/mainfile to run the correct server
	r.HandleFunc("/mainpage", handlers.IndexPage)
	// Create a new handler, make this handler return text saying "Hello from Jazz!, connect this handler to url /about"
	r.HandleFunc("/about", handlers.IndexPage2)
	// start server by putting the IP of the server usually its like "google.com" but if it's a local host then you can use port
	http.ListenAndServe(":3000", r)

}
