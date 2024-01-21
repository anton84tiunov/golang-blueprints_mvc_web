package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	config "../../internal/config"
	about "../blueprints/about"
	auth "../blueprints/auth"
	home "../blueprints/home"
	crud_user "../database/crud/user"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/base/static/").Handler(http.StripPrefix("/base/static/", http.FileServer(http.Dir("internal/app/static"))))
	router.PathPrefix("/favicon.ico").Handler(http.FileServer(http.Dir("internal/app/static/icons")))
	// router.Handle("/", http.RedirectHandler("/home", http.StatusSeeOther))

	router.PathPrefix("/about").Handler(about.AboutRoutes())
	router.PathPrefix("/auth").Handler(auth.AuthRoutes())

	router.PathPrefix("/").Handler(home.HomeRoutes())

	return router

}

func Srver(conf config.Config) error {
	router := SetupRoutes()

	// http.Handle("/", router)
	address := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	return http.ListenAndServe(address, router)
}

func Run() {
	conf := config.ReadConfig()
	fmt.Println(crud_user.CreateTable())

	// services.SendMessageMail("anton.tiunov.84.07@gmail.com", "test message", "Подтверждение почты")

	server := Srver(conf)
	fmt.Println(server)
}
