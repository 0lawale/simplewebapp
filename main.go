package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type HelloResp struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `<html><head><title>Simple Go philip's Webapp</title></head><body>
	<h1>Welcome to olawale's SimpleGoApp</h1>
	<p>This is a demo web app for a CI/CD pipeline project.</p>
	<p><a href="/api/hello">API: /api/hello</a> | <a href="/health">Health</a></p>
	</body></html>`
	t := template.Must(template.New("home").Parse(tmpl))
	_ = t.Execute(w, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp := HelloResp{
		Message: "Hello from SimpleGoApp",
		Time:    time.Now().UTC().Format(time.RFC3339),
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/api/hello", helloHandler)
	mux.HandleFunc("/health", healthHandler)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("starting webapp on :%s", port)
	log.Fatal(srv.ListenAndServe())
}
