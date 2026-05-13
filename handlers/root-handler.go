package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// valids URL path
	if r.URL.Path != "/" {
		// w.Write([]byte("page not found"))
		http.NotFound(w, r)
		return
	}

	// allows only GET
	if r.Method != http.MethodGet {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// Parse the template file
	temp, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Println("Template parsing error", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, nil)
	if err != nil {
		fmt.Println("Template execution error:", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
