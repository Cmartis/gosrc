//quick http handler - lite web

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the a http handler!")
	})
	fs := http.FileServer(http.Dir("statc/"))
	http.Handle("/static/", http.StripPrefix("static/", fs))
	http.ListenAndServe(":80", nil)
}
