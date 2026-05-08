package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type TemplateData struct {
	Result string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// w.Write([]byte("page not found"))
		http.NotFound(w, r)
		return
	}
	temp, _ := template.ParseFiles("./templates/index.html")
	temp.Execute(w, nil)

}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("In ascii art handler", r.Form)
	fmt.Println("Expected post", r.PostForm)

	inputText := r.Form.Get("inputText")
	banner := r.Form.Get("banner")

	result := asciiArt(inputText, banner)

	fmt.Println(result)

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, &TemplateData{Result: result})
}

func main() {

	http.HandleFunc("GET /", rootHandler)
	http.HandleFunc("POST /ascii-art", AsciiArtHandler)

	startRes := asciiArt("Our Server", "standard")
	fmt.Print(startRes)
	fmt.Println()

	fmt.Println("server starting on port 4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("error starting server:", err)
	}
}
