package main

import (
	asciiartcode "ascii-art-web/ascii-art-code"
	"ascii-art-web/handlers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("GET /", handlers.RootHandler)
	http.HandleFunc("POST /ascii-art", handlers.AsciiArtHandler)

	startRes, _ := asciiartcode.AsciiArt("Our Server", "standard")
	fmt.Print(startRes)
	fmt.Println()

	fmt.Println("server starting on port 4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("error starting server:", err)
	}
}
