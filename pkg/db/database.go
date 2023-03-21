package db

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nsrvel/golang-example/config"
)

func NewDBConnection(cfg *config.DatabaseAccount) *sqlx.DB {

	db, err := sqlx.Connect(cfg.ServerType, cfg.DriverSource)
	if err != nil {
		log.Fatalf("failed to connect database, err: %v", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime * time.Minute)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime * time.Minute)
	if err = db.Ping(); err != nil {
		log.Fatalf("failed to connect database, err: %v", err)
	}

	fmt.Println("Connection opened to database")
	return db
}
