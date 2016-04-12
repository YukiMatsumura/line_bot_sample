package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		_ = fmt.Errorf("$PORT not set")
		return
	}
	fmt.Println("PORT:" + port)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
