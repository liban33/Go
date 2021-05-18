package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","text/html" )
	
	
	fmt.Print("someone visited the")
	fmt.Fprint(w, "<h1>welcome to my new page1</h1>")

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":7000", nil)

}
