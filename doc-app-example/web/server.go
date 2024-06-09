// This package contains web server structures and functions responsible for handling HTTP application.
package web

import (
	"fmt"
	"net/http"
)

// Create an application web server mux with routes established
func BuildServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", HealthCheckEndpoint)
	mux.HandleFunc("GET /cards", CardGeneratorEndpoint)

	return mux
}

// Listening web server on port 4000
func ListenAndServe(mux *http.ServeMux) {
	fmt.Println("✅ The sever is listening at port 4000")
	http.ListenAndServe("localhost:4000", mux)
}
