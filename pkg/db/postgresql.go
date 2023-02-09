package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

func Connect() (db *gorm.DB, err error, connection *sql.DB) {

	db, err = gorm.Open(postgres.New(
		postgres.Config{DSN: os.Getenv("DSN"),
			PreferSimpleProtocol: true},
	),
		&gorm.Config{},
	)
	if err != nil {
		return
	}
	// logs
	db = db.Debug()
	connection, err = db.DB()
	if err != nil {
		return
	}
	connection.SetConnMaxLifetime(time.Duration(1) * time.Hour)
	connection.SetMaxIdleConns(15)
	connection.SetMaxOpenConns(15)
	fmt.Println("DB connected: ", os.Getenv("DSN"))
	return
}
