package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// w.Write([]byte("page not found"))
		http.NotFound(w, r)
		return
	}
	temp, _ := template.ParseFiles("./static/index.html")
	temp.Execute(w, nil)

}

func main() {

	http.HandleFunc("Get /", rootHandler)

	startRes := asciiArt("Our Server", "standard")
	fmt.Print(startRes)
	fmt.Println()

	fmt.Println("server starting on port 4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("error starting server:", err)

	}
}
