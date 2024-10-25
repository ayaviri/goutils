package http

import "net/http"

func StripQueryString(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.RawQuery = ""

		next.ServeHTTP(w, r)
	})
}
