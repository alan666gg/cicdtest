package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello cicd demo")
	})
	fmt.Println("listening on :8080")
	_ = http.ListenAndServe(":8080", nil)
}
