package middleware

import "net/http"

// CORS wraps an http.Handler to inject security headers
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight and HEAD requests directly
		if r.Method == http.MethodOptions || r.Method == http.MethodHead {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Proceed to the actual endpoint
		next.ServeHTTP(w, r)
	})
}
