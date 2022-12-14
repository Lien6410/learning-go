package main

import (
	"fmt"
	"net/http"

	"github.com/Lien6410/learning-go/web-cache1/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port " + portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
