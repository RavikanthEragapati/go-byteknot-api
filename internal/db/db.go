package db

import (
	"byteknot-api/configs"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLConnection(config *configs.Config) (*sql.DB, error) {
	log.Println("Establishing DB Connection...")
	conf := mysql.Config{
		User:   config.Database.User,
		Passwd: config.Database.Pass,
		Net:    config.Database.Net,
		Addr:   config.Database.Addr,
	}
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Set connection pool properties
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	// Ping the database to verify the connection is active
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}
