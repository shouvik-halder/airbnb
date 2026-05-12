package dbconfig

import (
	"AuthenticationService/config"
	"AuthenticationService/config/logger"
	"database/sql"
	"fmt"
	"os"
	"strings"

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

func SeedDB() error {
	content, err := os.ReadFile("db/seeder/seed_admin_user.sql")
	if err != nil {
		logger.Log.Error().Err(err).Msg("error reading seed file")
		return err
	}

	for statement := range strings.SplitSeq(string(content), ";") {
		statement = strings.TrimSpace(statement)

		if statement == "" {
			continue
		}

		if _, err := db.Exec(statement); err != nil {
			return fmt.Errorf("failed to execute seed statement: %w", err)
		}
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}
