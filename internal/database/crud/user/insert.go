package user

import (
	"fmt"
	"time"

	connect "../../connect"
)

func Insert_user(Name string, Surname string, Birthday time.Time, Email string, Phone string, Login string, Password []byte) error {

	db, err := connect.Connect()
	if err != nil {
		// panic(err_con)
		return err
	}

	defer db.Close()
	query := fmt.Sprintf("INSERT INTO `users` (`Name`, `Surname`, `Birthday`, `Email`,`Phone`, `Login`, `Password`) VALUES ('%s','%s','%s','%s','%s','%s','%s') ", Name, Surname, Birthday.Format("2006-01-02"), Email, Phone, Login, Password)
	insert, err_insert := db.Query(query)

	if err_insert != nil {
		fmt.Println("err_insert", err_insert)
		// panic(err_insert)
		return err_insert
	}
	defer insert.Close()
	return err_insert
}
