package main

import (
	"fmt"
	"log"
	"myapp/pkg/handlers"
	"net/http"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println(fmt.Sprintf("Server starts on http://127.0.0.1%s...", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
