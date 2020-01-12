package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

/*
 * database.Connect
 * SQLite NOT SUPPORTED
 * You must have .env file in your project
 * that have keys below
 */

// DBConnect to database
func Connect() *gorm.DB {
	var dialect = os.Getenv("DB_CONNECTION") // mysql | mssql | postgres
	var host = os.Getenv("DB_HOST")
	var port = os.Getenv("DB_PORT")
	var name = os.Getenv("DB_DATABASE")
	var username = os.Getenv("DB_USERNAME")
	var password = os.Getenv("DB_PASSWORD")

	connectionString := ""
	if dialect == "mysql" {
		connectionString = username + ":" + password + "@(" + host + ":" + port + ")/" + name + "?charset=utf8&parseTime=True&loc=Local"
	} else if dialect == "mssql" {
		connectionString = "sqlserver://" + username + ":" + password + "@" + host + ":" + port + "?database=" + name
	} else if dialect == "postgres" {
		connectionString = "host=" + host + " port=" + port + " user=" + username + " dbname=" + name + " password=" + password + " sslmode=disable"
	}

	database, err := gorm.Open(dialect, connectionString)
	if err != nil {
		log.Println("\nUnable to connect to database ...")
		log.Println("DETAILS: ")
		log.Println(err)
		os.Exit(1)
	}
	return database
}
