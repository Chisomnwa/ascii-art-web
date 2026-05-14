package handlers

import (
	asciiartcode "ascii-art-web/ascii-art-code"
	"fmt"
	"net/http"
	"text/template"
)

type TemplateData struct {
	Result string
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// validates the http post method
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	// validates url path
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Println("Form Parsing Error:", err)
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	// fmt.Println("In ascii art handler", r.Form)
	// fmt.Println("Expected post", r.PostForm)

	// get the form values: input text and banner
	inputText := r.Form.Get("inputText")
	banner := r.Form.Get("banner")

	if inputText == "" {
		http.Error(w, "404 Bad Request: input cannot be empty", http.StatusBadRequest)
		fmt.Println("")
		return
	}

	validBanner := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}

	if !validBanner[banner] {
		http.Error(w, "400 Bad Request: invalid banner selecton", http.StatusBadRequest)
		return
	}

	result, err := asciiartcode.AsciiArt(inputText, banner)
	if err != nil {
		fmt.Println("ascii-art generation error:", err)
		http.Error(w, "400 Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(result)

	// load the templates
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Println("Template parsing error:", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// inform the browser it is an html content
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = tmpl.Execute(w, &TemplateData{Result: result})
	if err != nil {
		fmt.Println("Template execution error", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
