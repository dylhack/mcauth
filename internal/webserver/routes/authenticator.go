package routes

import (
	"net/http"
)

// these makes sure every HTTP request to this web-server is authorized.
func (server *Server) authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")

		if len(token) == 0 {
			failAuth(res)
			return
		}

		if server.Config.Token == token {
			next.ServeHTTP(res, req)
		} else {
			failAuth(res)
			return
		}
	})
}

func failAuth(res http.ResponseWriter) {
	res.WriteHeader(http.StatusUnauthorized)
}
