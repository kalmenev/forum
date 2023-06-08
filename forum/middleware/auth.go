package middleware

import (
	"log"
	"net/http"
	"time"

	"forum/config"
	"forum/database"
)

// Authentication check middleware.
func AuthCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userCookie, err := r.Cookie("auth")

		// TODO: change if checks to guardian pattern for more readability.
		if err != nil {
			log.Println(err.Error())
			http.Redirect(w, r, config.LOGIN_PAGE, http.StatusFound)
			return
		}

		var cookieExpirationDate int64
		if err := database.DB.QueryRow(`SELECT "expires" from Sessions where "session"=?;`, userCookie.String()).Scan(&cookieExpirationDate); err != nil {
			log.Println(err.Error())
			http.Redirect(w, r, config.LOGIN_PAGE, http.StatusFound)
			return
		}

		if time.Now().Unix() >= cookieExpirationDate {
			log.Println("Cookie expired.")
			http.Redirect(w, r, config.LOGIN_PAGE, http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
