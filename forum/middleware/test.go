package middleware

import (
	"log"
	"net/http"

	"forum/config"
	"forum/entity"
	"forum/pages"
)

func CheckSingUpForm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		nickname := r.FormValue("nickname")
		password := r.FormValue("password")
		email := r.FormValue("email")
		firstName := r.FormValue("first_name")
		lastName := r.FormValue("last_name")

		if nickname == "" || password == "" || email == "" || firstName == "" || lastName == "" {
			w.WriteHeader(http.StatusBadRequest)
			pages.ALL.ExecuteTemplate(w, config.TemplateErrorPage, entity.Error{
				StatusCode: http.StatusBadRequest,
				StatusText: http.StatusText(http.StatusBadRequest),
			})
			// return
		}
		log.Println("done")
	})
}
