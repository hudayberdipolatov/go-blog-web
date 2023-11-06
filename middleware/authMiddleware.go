package middleware

import (
	"github.com/hudayberdipolatov/go-blog-web/helpers/authHelpers"
	"log"
	"net/http"
)

type Middleware struct{}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := authHelpers.Store.Get(r, authHelpers.SESSION_ID)
		if len(session.Values) == 0 {
			log.Println("user no login")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			log.Println("user login success")
			next.ServeHTTP(w, r)
			return
		}

	}
}
