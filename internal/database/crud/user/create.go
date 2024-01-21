package user

import (
	// "database/sql"
	"fmt"

	config "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/config"
	connect "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/database/connect"
	services "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/services"

	_ "github.com/go-sql-driver/mysql"
)

func IsCreatedTable() bool {
	db, err := connect.Connect()
	if err != nil {
		services.L.Warn(err)
		return false
	}
	defer db.Close()
	query := fmt.Sprintf("SHOW TABLES FROM %s LIKE 'users';", config.GLOBAL_CONFIG.Database.Dbname)
	insert, err_iscreate := db.Query(query)
	if err_iscreate != nil {
		services.L.Warn(err_iscreate)
		return false
	}
	defer insert.Close()
	return true
}

func CreateTable() error {
	db, err := connect.Connect()
	if err != nil {
		services.L.Warn(err)
		return err
	}
	defer db.Close()

	query := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS `users` (" +
			"`Id` int NOT NULL auto_increment," +
			"`Name` varchar(20) NOT NULL, " +
			"`Surname` varchar(20) NOT NULL, " +
			"`Birthday`  date NOT NULL, " +
			"`Email` varchar(50) UNIQUE, " +
			"`Phone`  varchar(15) NOT NULL, " +
			"`Login` varchar(20) UNIQUE, " +
			"`Password` blob NOT NULL, " +
			"PRIMARY KEY  (`id`)" +
			" );")
	create, err_create := db.Query(query)
	if err_create != nil {
		services.L.Warn(err_create)
	}
	defer create.Close()
	return err_create
}

// func Insert_user(db *sql.DB, err_con error) error {
// 	defer db.Close()
// 	if err_con == nil {
// 		// panic(err_con)
// 		return err_con
// 	}

// 	query := fmt.Sprintf("INSERT INTO `users` (`name`, `age`) VALUES ('%s', %d) ", "jfjfjfdjf", 14345)
// 	insert, err_insert := db.Query(query)

// 	if err_insert != nil {
// 		fmt.Println(err_insert)
// 		panic(err_insert)
// 	}
// 	defer insert.Close()
// 	return err_insert
// }
