package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	about "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/blueprints/about"
	auth "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/blueprints/auth"
	home "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/blueprints/home"
	config "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/config"
	crud_user "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/database/crud/user"
	services "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/services"
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
	address := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	if conf.Server.Debug {
		fmt.Printf("server run %s\n", address)
	}
	return http.ListenAndServe(address, router)
}

func Run() {
	conf := config.ReadConfig()
	isCreate := crud_user.CreateTable()
	if isCreate != nil {
		services.L.Warn(isCreate)
	}
	// services.SendMessageMail("anton.tiunov.84.07@gmail.com", "test message", "Подтверждение почты")
	server := Srver(conf)
	if server != nil {
		services.L.Warn(server)
	}
}
