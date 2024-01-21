package connect

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	config "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/config"
)

func Connect() (*sql.DB, error) {
	conf := config.GLOBAL_CONFIG
	db, err := sql.Open(conf.Database.Db, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.Dbname))
	return db, err
}
