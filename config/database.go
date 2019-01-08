package config

import (
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"           // mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgresql driver
	"github.com/joho/godotenv"
)

// Connection ..
type Connection struct {
	Driver     string
	DataSource string
	LogMode    bool
}

// Serve ..
type Serve struct {
	Port string
}

var confConnection *Connection
var confServe *Serve

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	drive := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	// select datasource
	var dataSource string
	switch drive {
	case "mysql":
		dataSource = userName + ":" + password +
			"@tcp(" + dbHost + ":" + dbPort + ")/" + database +
			"?charset=utf8&parseTime=True&loc=Local"
	case "postgres":
		dataSource = "host=" + dbHost +
			" port=" + dbPort +
			" user=" + userName +
			" dbname=" + database +
			" password=" + password +
			" sslmode=disable"
	}

	logMode, _ := strconv.ParseBool(os.Getenv("DB_LOGMODE"))
	confConnection = &Connection{
		Driver:     drive,
		DataSource: dataSource,
		LogMode:    logMode,
	}

	servePort := os.Getenv("SERVE_PORT")
	confServe = &Serve{
		Port: servePort,
	}
}

// GetConnection ..
func GetConnection() *Connection {
	return confConnection
}

// GetServe ..
func GetServe() *Serve {
	return confServe
}
