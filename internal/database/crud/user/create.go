package user

import (
	// "database/sql"
	"fmt"

	connect "../../connect"

	config "../../../config"
	_ "github.com/go-sql-driver/mysql"
)

func IsCreatedTable() bool {
	db, err := connect.Connect()
	if err != nil {
		return false
	}
	defer db.Close()
	query := fmt.Sprintf("SHOW TABLES FROM %s LIKE 'users';", config.GLOBAL_CONFIG.Database.Dbname)
	insert, err_iscreate := db.Query(query)
	if err_iscreate != nil {
		return false
	}
	defer insert.Close()
	return true
}

func CreateTable() error {
	db, err := connect.Connect()
	if err != nil {
		// panic(err)
		return err
	}
	defer db.Close()

	query := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS `users` (" +
			"`Id` int NOT NULL auto_increment," +
			"`Name` varchar(20) NOT NULL, " +
			"`Surname` varchar(20) NOT NULL, " +
			"`Birthday`  date NOT NULL, " +
			"`Email` varchar(50)  NOT NULL, " +
			"`Phone`  varchar(15) NOT NULL, " +
			"`Login` varchar(20), " +
			"`Password` blob NOT NULL, " +
			"UNIQUE (`Login`), " +
			"PRIMARY KEY  (`id`)" +
			" );")
	create, err_create := db.Query(query)
	fmt.Println("create", create.NextResultSet())
	if err_create != nil {
		fmt.Println("err_create", err_create)
		panic(err_create)
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
