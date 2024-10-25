package http

import (
	"fmt"
	"net/http"
)

func StripQueryString(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.RawQuery = ""

		fmt.Println(r.URL)

		next.ServeHTTP(w, r)
	})
}
