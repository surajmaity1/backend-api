package config

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func dbConnect() bool {
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = AppConfig.MYSQL_USER
	mysqlConfig.Passwd = AppConfig.MYSQL_PASS
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = AppConfig.MYSQL_HOST
	mysqlConfig.DBName = AppConfig.MYSQL_DATABASE

	var err error
	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		logrus.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		logrus.Fatal(pingErr)
	}
	
	return true
}
