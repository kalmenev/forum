package endpoint

import (
	"net/http"

	"forum/config"
	"forum/entity"
	"forum/pages"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Allow", http.MethodGet)
		pages.ALL.ExecuteTemplate(w, config.TemplateErrorPage, entity.Error{
			StatusCode: http.StatusMethodNotAllowed,
			StatusText: http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	if r.URL.Path == config.MAIN_PAGE {
		pages.ALL.ExecuteTemplate(w, "index.html", nil)
		return
	}
}
