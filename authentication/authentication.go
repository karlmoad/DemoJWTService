package authentication

import (
	"github.com/gorilla/context"
	"net/http"
)

func AuthenticationHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate token
		// if not valid
		// TODO - set header status to 401 amd discontinue request process
		// else
		// TODO - implement, this is stand in for now to allow testing
		context.Set(r, "USER", "user1")
		next(w, r)
	})
}
