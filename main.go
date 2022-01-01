package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world")
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
		if err != nil {
			log.Println(err)
		}
	})

	log.Println("Server starts on port 8080...")
	_ = http.ListenAndServe(":8080", nil)
}
