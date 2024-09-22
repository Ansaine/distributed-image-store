package main

import (
	"dkv/routes"
	"fmt"
	"net/http"
)

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Allowed headers

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	port := ":8080"

	// Create a new ServeMux and wrap it with the CORS middleware
	mux := http.NewServeMux()
	mux.HandleFunc("/set", routes.Set)
	mux.HandleFunc("/get", routes.Get)

	// Wrap the mux with the CORS middleware
	fmt.Println("server starting on port", port)
	err := http.ListenAndServe(port, enableCORS(mux))
	if err != nil {
		fmt.Println("error starting server:", err)
	}
}
