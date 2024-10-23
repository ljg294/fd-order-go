package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// SetupDB initializes the MySQL database connection
func SetupDB() (*sql.DB, error) {
	// Use your own credentials and database details
	dsn := "root:root@tcp(localhost:22201)/local_order?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verify connection to the database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")
	return db, nil
}
