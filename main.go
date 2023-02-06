package main

import (
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Hello, World! </b>")
	fmt.Fprintf(w, "You are at %s", r.URL.Path)
}

func main() {

	http.HandleFunc("/", greet)
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)

}
