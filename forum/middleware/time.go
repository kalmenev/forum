package middleware

import (
	"log"
	"net/http"
	"time"
)

func MesureRequestTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		log.Println(elapsed)
	})
}
