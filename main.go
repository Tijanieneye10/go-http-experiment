package main

import (
	"net/http"
	_"github.com/russross/blackfriday"
)


func main() {
	
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir(".")))

	// http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}


func GenerateMarkdown(w http.ResponseWriter, r *http.Request) {
	// markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))	
	w.Write([]byte(r.URL.Query().Get("q")))
	// w.Write(markdown)
}
