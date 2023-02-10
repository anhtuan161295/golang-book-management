package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config "example.com/openweather/config"
)

var dbConn *gorm.DB

func GetDbConnection() (db *gorm.DB) {
	// reuse database connection
	if dbConn != nil {
		return dbConn
	}

	databaseConfig := config.MyDatabaseConfig

	//dbDriver := "mysql"
	dbUrl := databaseConfig.Url
	dbPort := databaseConfig.Port
	dbUser := databaseConfig.Username
	dbPass := databaseConfig.Password
	dbName := databaseConfig.DatabaseName

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbUrl, dbPort, dbName)
	// connString := dbUser + ":" + dbPass + "@tcp(" + dbUrl + ":" + dbPort + ")/" + dbName

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// set connection pool
	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	dbConn := db

	return dbConn
}
