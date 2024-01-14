package user

import (
	"fmt"
	"log"
	"time"

	models "../../../models"
	connect "../../connect"

	_ "github.com/go-sql-driver/mysql"
)

func Select_user(Login string) models.User {

	user := models.User{}
	datestr := ""
	db, err := connect.Connect()
	if err != nil {
		// panic(err_con)
		return user
	}
	defer db.Close()
	query := fmt.Sprintf("Select * from `users` WHERE `Login`='%s';", Login)
	sel, err_select := db.Query(query)

	if err_select != nil {
		fmt.Println(err_select)
		panic(err_select)
	}
	for sel.Next() {
		err := sel.Scan(&user.Id, &user.Name, &user.Surname, &datestr, &user.Email, &user.Phone, &user.Login, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(user)
	}
	user.Birthday, _ = time.Parse("2006-01-02", datestr)
	// fmt.Println(sel.Columns())
	defer sel.Close()
	return user
}
