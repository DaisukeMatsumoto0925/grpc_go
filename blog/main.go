package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)

	fmt.Println("start server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
