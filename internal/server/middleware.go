package server

import (
	"fmt"
	"net/http"
	"time"
)

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(w, r)
		end := time.Since(begin)
		s.Logger.Info(
			fmt.Sprintf("%s %s %s %d [%dms] %s",
				r.RemoteAddr,
				r.Method,
				r.RequestURI,
				http.StatusOK,
				end.Microseconds(),
				r.UserAgent()),
		)
	})
}
