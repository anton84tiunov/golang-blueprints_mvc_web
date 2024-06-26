package home

import (
	"net/http"

	home_handle "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/blueprints/home/handlers"

	"github.com/gorilla/mux"
)

func HomeRoutes() http.Handler {
	router := mux.NewRouter()

	router.PathPrefix("/home/static/").Handler(http.StripPrefix("/home/static/", http.FileServer(http.Dir("internal/blueprints/home/static"))))

	router.HandleFunc("/home/", home_handle.HomeHandler)
	router.HandleFunc("/", home_handle.HomeHandler)

	return router
}
