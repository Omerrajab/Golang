// package dbHelpers

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// var db *sql.DB

// func InitDB() error {
// 	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	var err error
// 	db, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		return err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Connected to the PostgreSQL database")
// 	return nil
// }
// func GetDB() *sql.DB {
// 	return db
// }
// helpers/dbHelpers.go

package dbHelpers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the PostgreSQL database connection

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "GoDB"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	//connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
