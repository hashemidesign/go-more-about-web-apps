package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("about page. 2 + 2 = %d", sum))
}

func AddValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x float32 = 100.0
	var y float32 = 0.0
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f = %f", x, y, f))
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	return x / y, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	log.Println(fmt.Sprintf("Server starts on http://127.0.0.1%s...", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
