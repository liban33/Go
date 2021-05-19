package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my website</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "Get in touch, send an email to <a href=\"mailto:liban_jama@hotmail.com\">liban_jama@hotmail.com</>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>we could not find your page</h1><p>please email us if you keep being sent to an invalid page.</p>")

	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my website</h1>")
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/Html")
	fmt.Fprint(w, "Get in touch, send an email to <a href=\"mailto:liban_jama@hotmail.com\">liban_jama@hotmail.com</>.")
}
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html")
	fmt.Fprint(w, "<h1>frequently asked Questions?</h1><p> Here is a list of questions that our users commonly ask.</p>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	//http.HandleFunc("/contact", handlerFunc)
	http.ListenAndServe(":7000", r)
}

