package main

import (
	"database/sql"
	"log"

	"github.com/emirrcaglar/go-restapi/cmd/api"
	"github.com/emirrcaglar/go-restapi/config"
	"github.com/emirrcaglar/go-restapi/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal("error loading config: ", err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal("error running api server: ", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("error pinging database: ", err)
	}
	log.Println("db: successfully connected.")
}
