package endpoint

import (
	"log"
	"net/http"

	"forum/pages"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("incoming request: %s | %s\n", r.URL.Path, r.Method)

	switch r.Method {
	case http.MethodGet:
		pages.ALL.ExecuteTemplate(w, "sign-up.html", nil)
	case http.MethodPost:
		nickname := r.FormValue("nickname")
		password := r.FormValue("password")
		email := r.FormValue("email")
		firstName := r.FormValue("first_name")
		lastName := r.FormValue("last_name")
		log.Println("=============================",nickname, password, email, firstName, lastName)

		
		// TODO: CHECK IF THE PROVIDED DATA IS NOT EMPTY.
		// TODO: WRITE IT INTO THE DATABASE.

		// TODO: CHECK IF REDIRECTION WORKS CORRECTLY.
		http.Redirect(w, r, "/login", http.StatusFound)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

	}
}
