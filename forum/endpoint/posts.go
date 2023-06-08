package endpoint

import (
	"log"
	"net/http"

	"forum/pages"
)

func PostPageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("incoming request: %s | %s\n", r.URL.Path, r.Method)
	switch r.Method {
	case http.MethodGet:
		pages.ALL.ExecuteTemplate(w, "allposts.html", nil)
	case http.MethodPost:
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	pages.ALL.ExecuteTemplate(w, "createpost.html", nil)
}
