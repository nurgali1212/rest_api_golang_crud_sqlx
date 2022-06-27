package database

import (
	"fmt"
	"log"
	"test_project/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDatabase() *sqlx.DB {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	psqlconn := fmt.Sprintf("user = %s password = %s port = %s dbname = %s sslmode=disable", config.Username, config.Password, config.Port, config.DBName)
	db, err := sqlx.Connect("postgres", psqlconn)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("sksdm")
	}

	return db
}
