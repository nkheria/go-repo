package main

import (
	"fmt"
	"net/http"

	"example.com/go-demo-1/pkg/handlers"
)

const portNumber = ":8081"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Application running on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
