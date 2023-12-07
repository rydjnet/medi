package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

type PgConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}
type EnvironmentConfigProvider struct{}

func (e EnvironmentConfigProvider) GetPostgresConfig() PgConfig {
	host := os.Getenv("PG_HOST")
	portStr := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASS")
	dbname := os.Getenv("PG_DBNAME")
	if host == "" {
		host = "localhost"
	}
	port := 5432
	if portStr != "" {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatal(err)
		}
		port = p
	}
	if user == "" {
		user = "postgres"
	}
	if password == "" {
		log.Fatal("PG_PASS is not defined")
	}
	if dbname == "" {
		log.Fatal("PG_DBNAME is not defined")
	}
	return PgConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Dbname:   dbname,
	}
}

var DB *sql.DB

func ConnectDB() {
	configProvider := EnvironmentConfigProvider{}

	pConf := configProvider.GetPostgresConfig()
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pConf.Host, pConf.Port, pConf.User, pConf.Password, pConf.Dbname)

	// open database
	d, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	// check db
	err = d.Ping()
	if err != nil {
		log.Fatal(err)
	}
	DB = d
	fmt.Println("Connected!")
}
