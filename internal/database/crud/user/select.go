package user

import (
	"fmt"
	"time"

	connect "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/database/connect"
	models "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/models"
	services "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/services"

	_ "github.com/go-sql-driver/mysql"
)

func Select_user(Login string) models.User {

	user := models.User{}
	datestr := ""
	db, err := connect.Connect()
	if err != nil {
		services.L.Warn(err)
		return user
	}
	defer db.Close()
	query := fmt.Sprintf("Select * from `users` WHERE `Login`='%s';", Login)
	sel, err_select := db.Query(query)

	if err_select != nil {
		services.L.Warn(err_select)
	}
	for sel.Next() {
		err := sel.Scan(&user.Id, &user.Name, &user.Surname, &datestr, &user.Email, &user.Phone, &user.Login, &user.Password)
		if err != nil {
			services.L.Warn(err)
		}
	}
	user.Birthday, _ = time.Parse("2006-01-02", datestr)

	defer sel.Close()
	return user
}
