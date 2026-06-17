package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
