package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	DBConnect *gorm.DB
)

func Init() {
	var dialect = os.Getenv("DB_CONNECTION") // mysql | sqlserver | postgres
	var host = os.Getenv("DB_HOST")
	var port = os.Getenv("DB_PORT")
	var name = os.Getenv("DB_DATABASE")
	var username = os.Getenv("DB_USERNAME")
	var password = os.Getenv("DB_PASSWORD")
	// Other optional parameters
	var timezone = os.Getenv("TIME_ZONE")
	// for mysql only
	var charset = os.Getenv("CHAR_SET")
	var parseTime = os.Getenv("PARSE_TIME")
	// for postgres only
	var sslmode = os.Getenv("SSL_MODE")
	// error variable
	var err error

	// Set default value for optional parameters
	if timezone == "" {
		timezone = "Africa/Lagos"
	}
	if charset == "" {
		charset = "utf8mb4"
	}
	if parseTime == "" {
		parseTime = "True"
	}
	if sslmode == "" {
		sslmode = "disable"
	}

	// set connection for database
	if dialect == "mysql" {
		DBConnect, err = mysqlConnection(username, password, host, port, name, charset, parseTime, timezone)
	} else if dialect == "sqlserver" {
		DBConnect, err = sqlserverConnection(username, password, host, port, name)
	} else if dialect == "postgres" {
		DBConnect, err = postgresConnection(host, username, password, name, port, sslmode, timezone)
	}

	if err != nil {
		log.Printf("Unable to connect to database ...\nDETAILS: \n%v", err)
		os.Exit(1)
	}
}

// mysqlConnection connect to mysql or maria database
func mysqlConnection(username, password, host, port, name, charset, parseTime, timezone string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", username, password, host, port, name, charset, parseTime, timezone)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn, // data source name
		DefaultStringSize: 255, // default size for string fields
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	return db, err
}

// sqlserverConnection connect to microsoft sqlserver database
func sqlserverConnection(username, password, host, port, name string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", username, password, host, port, name) // sqlserver://username:password@host:port?database=name

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	return db, err
}

// postgresConnection connect to postgres database
func postgresConnection(host, username, password, name, port, sslmode, timezone string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, username, password, name, port, sslmode, timezone) // data source name

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,  // data source name
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	return db, err
}
