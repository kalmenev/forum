package endpoint

import (
	"net/http"

	"forum/config"
	"forum/database"
	"forum/pages"

	"golang.org/x/crypto/bcrypt"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		pages.ALL.ExecuteTemplate(w, "sign-in.html", nil)
	case http.MethodPost:
		login := r.FormValue("login")
		password := r.FormValue("password")
		println(login, password)

		// check active session, login and password is correct.
		query := `SELECT EXISTS (
			SELECT 1
			FROM Sessions
			INNER JOIN Users ON Sessions.user_id = Users.user_id
			WHERE Sessions.user_id = ?
				AND Sessions.expires > strftime('%s', 'now')
				AND Users.password = ?
		) AS active_session`
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			// TODO: Check possible redundant return after http.Error() call.
			return
		}

		row := database.DB.QueryRow(query, login, string(hash))
		if row.Err() != nil {
			// TODO
		}

		// TODO: read result from query

		http.Redirect(w, r, config.TemplateIndexPage, http.StatusFound)
		// TODO: CHECK THE GIVEN FIELDS.
		// TODO: SAVE TO THE DATABASE.
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
