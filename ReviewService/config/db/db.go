package dbConfig

import (
	"ReviewService/config"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SetupDB(cfg *config.Config) error {
	dbcfg := mysql.NewConfig()
	dbcfg.User = cfg.DB.DBUSER
	dbcfg.Passwd = cfg.DB.DBPASS
	dbcfg.Addr = cfg.DB.DBADDR
	dbcfg.Net = cfg.DB.DBNET
	dbcfg.DBName = cfg.DB.DBNAME
	dbcfg.ParseTime = true
	var err error
	db, err = sql.Open("mysql", dbcfg.FormatDSN())
	if err != nil {

		fmt.Println("Error connecting to db", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging db", err)
		return err
	}
	fmt.Println("DB connected")

	return nil
}

func GetDB() *sql.DB {
	return db
}
