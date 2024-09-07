package main

import (
    "html/template"
    "net/http"
)

var templates = template.Must(template.ParseFiles("templates/home.html", "templates/about.html"))

func home(write http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(write, "home.html", nil)
    if err != nil {
        http.Error(write, err.Error(), http.StatusInternalServerError)
    }
}

func about(write http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(write, "about.html", nil)
    if err != nil {
        http.Error(write, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/about", about)
    http.ListenAndServe(":8080", nil)
}
