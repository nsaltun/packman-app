package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	apiBaseURL string
	templates  *template.Template
)

type PageData struct {
	ApiBaseURL string
}

func main() {
	// Load API URL from environment variable
	apiBaseURL = os.Getenv("API_BASE_URL")
	if apiBaseURL == "" {
		apiBaseURL = "http://localhost:8081" // Default for local development
	}
	log.Printf("Using API: %s\n", apiBaseURL)

	// Load templates
	templatePath := filepath.Join("web", "templates", "*.html")
	var err error
	templates, err = template.ParseGlob(templatePath)
	if err != nil {
		log.Fatal("Error loading templates:", err)
	}

	// Routes
	http.HandleFunc("/", serveHome)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default for local development
	}

	log.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := PageData{
		ApiBaseURL: apiBaseURL,
	}

	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
