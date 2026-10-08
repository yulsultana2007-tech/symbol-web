package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Text string
	}{
		Text: "Eto napisal ya",
	}

	w.Header().Set("Contetn-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("test/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("server zapushen на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
