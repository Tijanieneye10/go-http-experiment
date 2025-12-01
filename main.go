package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)


func main() {

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir(".")))

	// http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}


func GenerateMarkdown(w http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))	
	// w.Write([]byte(r.URL.Query().Get("q")))
	w.Write(markdown)
}
