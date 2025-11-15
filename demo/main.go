package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/nyxstack/scalarui"
)

var ui *scalarui.ScalarUI

func main() {

	config := scalarui.NewConfig()
	config.URL = "http://localhost:8080/openapi.yaml"
	config.HotReloadURL = "http://localhost:8080/hot-reload"
	ui = scalarui.New(config)

	// cors on all routes
	handler := withCORS(http.DefaultServeMux)

	// Setup routes
	http.HandleFunc("/", RenderScalarUi)
	http.HandleFunc("/hot-reload", hotReloadHandler)
	http.HandleFunc("/openapi.yaml", yamlHandler)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}

var startTime = time.Now()

func hotReloadHandler(w http.ResponseWriter, r *http.Request) {
	// send a simple timestamp string
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")

	fmt.Fprintf(w, "%d", startTime.Unix())
}

// RenderScalarUi serves a simple hello world message
func RenderScalarUi(w http.ResponseWriter, r *http.Request) {

	str, err := ui.Render()
	if err != nil {
		http.Error(w, "Error rendering UI", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, str)
}

// yamlHandler serves YAML data
func yamlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-yaml; charset=utf-8")
	http.ServeFile(w, r, filepath.Join("demo/data", "openapi.yaml"))
}
