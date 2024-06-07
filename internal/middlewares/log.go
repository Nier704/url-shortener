package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("%s %s %s %s", r.Method, r.Host, r.URL, time.Now().Format(time.DateTime))
		next.ServeHTTP(w, r)
		log.Println(msg)
	})
}
