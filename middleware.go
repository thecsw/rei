package rei

import "net/http"

// BearerMiddleware is a middleware that checks for a bearer token.
func BearerMiddleware(token string) func(h http.Handler) http.Handler {
	// If the auth token is empty, return a no-op middleware.
	if len(token) < 1 {
		return NoopMiddleware()
	}

	// If the auth token is not empty, check that it matches the Authorization header.
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check that the Authorization header is present.
			if "Bearer "+token != r.Header.Get("Authorization") {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("unauthorized"))
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

// NoopMiddleware is a middleware that does nothing.
func NoopMiddleware() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		})
	}
}
