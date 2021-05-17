package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>welcome to my page1</h1>")

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":7000", nil)

}
