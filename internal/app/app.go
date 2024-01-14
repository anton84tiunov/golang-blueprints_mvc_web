package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	config "../../internal/config"
	services "../../internal/services"
	logprint "../../pkg/log_print"
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
	logprint.Print(conf.Database.Db)

	hash_pass, err := services.Hashing("passs123@")
	if err != nil {
		panic(err)
	}
	logprint.Print(string(hash_pass))

	// fmt.Println(crud_user.IsCreatedTable())
	fmt.Println(crud_user.CreateTable())
	err_ins := crud_user.Insert_user("aa", "aa", time.Now(), "<EMAIL>", "1234567", "log", hash_pass)
	fmt.Println(err_ins)
	usr := crud_user.Select_user("log")

	err = services.CompareHashAndPassword(usr.Password, "passs123@")
	if err != nil {
		// panic(err)
		logprint.Print("password is wrong")
	} else {
		logprint.Print("password is correct")
	}
	server := Srver(conf)
	fmt.Println(server)
}
