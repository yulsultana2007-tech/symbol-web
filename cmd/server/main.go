package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
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
	bannerName := banner + ".txt"

	lines, err := os.ReadFile(bannerName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	content := string(lines)

	bannerLines := strings.Split(content, "\n")

	fmt.Println(bannerName)

	result := generateArt(text, bannerLines)
	fmt.Fprint(w, result)
}

func generateArt(word string, lines []string) string {
	var builder strings.Builder

	for row := 0; row < 8; row++ {
		for _, symbol := range word {
			index := int(symbol) - 32
			start := index*9 + 1

			builder.WriteString(lines[start+row])

		}
		builder.WriteString("\n")

	}
	result := builder.String()
	return result
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
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server upal", err)
	}
}
