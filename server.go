package main

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templ.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", mux)
}

// ----- 05 using bootstrap
