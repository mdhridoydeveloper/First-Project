package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // 1:37:15 s

		next.ServeHTTP(w, r) // 10 s
		fmt.Println(r.Method, r.URL.Path, time.Since((start)))

	})
}
