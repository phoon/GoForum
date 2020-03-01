package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	Fields *fields
	_once  sync.Once
)

type fields struct {
	Mode          string //Gin mode [release/debug]
	Address       string //address to running
	BaseURL       string //base domain name
	GinLogFile    string //Gin framwork log file
	GinSessionKey string //Secret key for session
	Database      database
}

type database struct {
	Host     string
	Username string
	Password string
	DBName   string
}

func parse() {
	Fields = &fields{}
	if err := godotenv.Load(); err != nil {
		log.Fatalln("no .env configuration file found")
	}

	if Fields.Mode = os.Getenv("GIN_MODE_ENV"); Fields.Mode == "" {
		log.Fatalln("no GIN_MODE_ENV found")
	}
	if Fields.Address = os.Getenv("GIN_ADDRESS_ENV"); Fields.Address == "" {
		log.Fatalln("no GIN_ADDRESS_ENV found")
	}
	if Fields.BaseURL = os.Getenv("GIN_BASE_URL_ENV"); Fields.BaseURL == "" {
		log.Fatalln("no GIN_BASE_URL_ENV found")
	}
	if Fields.GinLogFile = os.Getenv("GIN_LOGFILE_ENV"); Fields.GinLogFile == "" {
		log.Fatalln("no GIN_LOGFILE_ENV found")
	}
	if Fields.GinSessionKey = os.Getenv("GIN_SESSION_KEY_ENV"); Fields.GinSessionKey == "" {
		log.Fatalln("no GIN_SESSION_KEY_ENV found")
	}
	if Fields.Database.Host = os.Getenv("DB_HOST_ENV"); Fields.Database.Host == "" {
		log.Fatalln("no DB_HOST_ENV found")
	}
	if Fields.Database.Username = os.Getenv("DB_USER_ENV"); Fields.Database.Username == "" {
		log.Fatalln("no DB_USER_ENV found")
	}
	if Fields.Database.Password = os.Getenv("DB_PASS_ENV"); Fields.Database.Password == "" {
		log.Fatalln("no DB_PASS_ENV found")
	}
	if Fields.Database.DBName = os.Getenv("DB_NAME_ENV"); Fields.Database.DBName == "" {
		log.Fatalln("no DB_NAME_ENV found")
	}
}

func Load() {
	_once.Do(parse)
}
