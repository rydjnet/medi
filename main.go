package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
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

func main() {
	configProvider := EnvironmentConfigProvider{}

	pConf := configProvider.GetPostgresConfig()
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pConf.Host, pConf.Port, pConf.User, pConf.Password, pConf.Dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
