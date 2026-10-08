package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func symbolArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if banner != "shadow" &&
		banner != "standard" &&
		banner != "thinkertoy" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, text, banner)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Text string
	}{
		Text: "Hello fron GO",
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/symbol-art", symbolArtHandler)
	fmt.Println("server zapushen на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
