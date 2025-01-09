package db

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"strconv"
)

func NewDB() *sqlx.DB {
	dbName := os.Getenv("PG_DB")
	dbHost := os.Getenv("PG_HOST")
	dbPort := os.Getenv("PG_PORT")
	dbUser := os.Getenv("PG_USER")
	dbPass := os.Getenv("PG_PASS")
	maxConns := os.Getenv("PG_MAX_CONNECTIONS")

	postgresConnStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPass)

	connConfig, err := pgx.ParseConfig(postgresConnStr)
	if err != nil {
		log.Panic(err)
	}

	connStr := stdlib.RegisterConnConfig(connConfig)

	conn, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		log.Panic(err)
	}

	if err := conn.Ping(); err != nil {
		log.Panic(err)
	} else {
		maxConns, err := strconv.Atoi(maxConns)
		if err != nil {
			conn.SetMaxOpenConns(5) // default 5 max connections in pool
		} else {
			conn.SetMaxOpenConns(maxConns)
		}
	}

	return conn
}
