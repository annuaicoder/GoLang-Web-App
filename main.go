package main

import (
	"fmt"
	"net/http"
)


func home(write http.ResponseWriter, r*http.Request) {
	fmt.Fprintf(write,"Hello World!")
}

func about(write http.ResponseWriter, r*http.Request) {
	fmt.Fprintf(write,"About us!")
}

func main() {
	http.HandleFunc("/",home)
	http.HandleFunc("/about",about)
	http.ListenAndServe(":8080",nil)
}