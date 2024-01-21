package auth

import (
	"net/http"

	auth_handl "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/blueprints/auth/handlers"

	"github.com/gorilla/mux"
)

func AuthRoutes() http.Handler {
	router := mux.NewRouter()

	router.PathPrefix("/auth/static/").Handler(http.StripPrefix("/auth/static/", http.FileServer(http.Dir("internal/blueprints/auth/static"))))

	router.HandleFunc("/auth/auth", auth_handl.AuthHandler)
	router.HandleFunc("/auth/reg", auth_handl.RegHandler)
	router.HandleFunc("/auth/reg_check", auth_handl.RegCheckHandler)

	return router
}
