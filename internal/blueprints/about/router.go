package about

import (
	"net/http"

	about_handl "./handlers"

	"github.com/gorilla/mux"
)

func AboutRoutes() http.Handler {
	router := mux.NewRouter()

	router.PathPrefix("/about/static/").Handler(http.StripPrefix("/about/static/", http.FileServer(http.Dir("internal/blueprints/about/static"))))

	router.HandleFunc("/about/", about_handl.AboutHandler)
	// router.HandleFunc("/", AboutHandler)
	return router
}
