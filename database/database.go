package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config "example.com/openweather/config"
)

func GetDbConnection() (db *gorm.DB) {
	databaseConfig := config.MyDatabaseConfig

	//dbDriver := "mysql"
	dbUrl := databaseConfig.Url
	dbPort := databaseConfig.Port
	dbUser := databaseConfig.Username
	dbPass := databaseConfig.Password
	dbName := databaseConfig.DatabaseName
	connString := dbUser + ":" + dbPass + "@tcp(" + dbUrl + ":" + dbPort + ")/" + dbName

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
