package pgdb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type ConfigDatabase struct {
	DriverName string
	Host       string
	Port       string
	Username   string
	Password   string
	DBName     string
	SSLMode    string
}

var DB = func() *sqlx.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	configDatabase := ConfigDatabase{
		DriverName: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		Username:   os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		SSLMode:    os.Getenv("DB_SSL_MODE"),
	}

	db, err := sqlx.Connect(
		configDatabase.DriverName,
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			configDatabase.Host,
			configDatabase.Port,
			configDatabase.Username,
			configDatabase.DBName,
			configDatabase.Password,
			configDatabase.SSLMode,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
