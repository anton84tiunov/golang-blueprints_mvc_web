package auth

import (
	"net/http"

	auth_handl "./handlers"

	"github.com/gorilla/mux"
)

func AuthRoutes() http.Handler {
	router := mux.NewRouter()

	router.PathPrefix("/auth/static/").Handler(http.StripPrefix("/auth/static/", http.FileServer(http.Dir("internal/blueprints/auth/static"))))

	router.HandleFunc("/auth/auth", auth_handl.AuthHandler)
	router.HandleFunc("/auth/reg", auth_handl.RegHandler)

	return router
}
