package user

import (
	"fmt"

	connect "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/database/connect"
	models "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/models"
	services "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/services"
)

func Existes_col_user(key string, val string) bool {
	user := models.User{}
	db, err := connect.Connect()
	if err != nil {
		services.L.Warn(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT  EXISTS(SELECT `Id` FROM `users` WHERE `%s`='%s');", key, val)
	sel, err_sel := db.Query(query)

	if err_sel != nil {
		services.L.Warn(err_sel)
	}
	for sel.Next() {
		err := sel.Scan(&user.Id)
		if err != nil {
			services.L.Warn(err)
		}
	}
	defer sel.Close()

	return user.Id > 0
}
