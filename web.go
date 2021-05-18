package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/"{
		fmt.Fprint(w, "<h1>Welcome to my website</h1>")
	} else if r.URL.Path == "/contact"{
		fmt.Fprint(w, "Get in touch, send an email to <a href=\"mailto:liban_jama@hotmail.com\">liban_jama@hotmail.com</>.")
		}
}
	
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":7000", nil)

}
