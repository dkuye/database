package database

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
		connectionString = username + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
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

	database.SetLogger(&GormLogger{})
	database.LogMode(true)

	return database
}

// Using logrus and gorm
// GormLogger is a custom logger for Gorm, making it use logrus.
type GormLogger struct{}

/* func (*GormLogger) Print(v ...interface{}) {
    if v[0] == "sql" {
        log.WithFields(log.Fields{"module": "gorm", "type": "sql"}).Print(v[3])
    }
    if v[0] == "log" {
        log.WithFields(log.Fields{"module": "gorm", "type": "log"}).Print(v[2])
    }
} */

/*
Nice!
Slight variation I made:
*/
// Print handles log events from Gorm for the custom logger.
func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		log.WithFields(
			log.Fields{
				"module":  "gorm",
				"type":    "sql",
				"rows":    v[5],
				"src_ref": v[1],
				"values":  v[4],
			},
		).Debug(v[3])
	case "log":
		log.WithFields(log.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
