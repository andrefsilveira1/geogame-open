package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/andrefsilveira1/LoadEnv"
	_ "github.com/microsoft/go-mssqldb"
)

func ConnectAzure() (*sql.DB, error) {
	var db *sql.DB
	var server = LoadEnv.LoadEnv("DB_AZURE_SERVER")
	var port = 1433
	var user = LoadEnv.LoadEnv("DB_AZURE_LOGIN")
	var password = LoadEnv.LoadEnv("DB_AZURE_PASSWORD")
	var database = LoadEnv.LoadEnv("DB_AZURE_DATABASE")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db, nil
}
