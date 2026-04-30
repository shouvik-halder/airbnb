package dbconfig

import (
	"AuthenticationService/config"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SetupDB(cfg *config.Config) error {

	dbcfg := mysql.NewConfig()
	dbcfg.User = cfg.DB.DBUSER
	dbcfg.Passwd = cfg.DB.DBPASS
	dbcfg.Net = cfg.DB.DBNET
	dbcfg.Addr = cfg.DB.DBADDR
	dbcfg.DBName = cfg.DB.DBNAME
	dbcfg.ParseTime = true

	var err error
	db, err = sql.Open("mysql", dbcfg.FormatDSN())
	if err != nil {
		fmt.Println("Issue while connecting to DB", err.Error())
		return err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Issue while pinging to DB", err.Error())
		return err
	}

	fmt.Println("DB Connected!")
	return nil
}

func GetDB() *sql.DB {
	return db
}
