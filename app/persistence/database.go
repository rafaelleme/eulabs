package persistence

import (
	"database/sql"
	"fmt"
	"github.com/subosito/gotenv"
	"os"
	"time"
)

func InitDB() (*sql.DB, error) {
	err := gotenv.Load()
	if err != nil {
		panic("Err on load envs")
	}

	var db *sql.DB

	dbHost := os.Getenv("MYSQL_DB_HOST")
	dbPort := os.Getenv("MYSQL_DB_PORT")
	dbName := os.Getenv("MYSQL_DB_NAME")
	dbUser := os.Getenv("MYSQL_DB_USER")
	dbPassword := os.Getenv("MYSQL_DB_PASSWORD")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	fmt.Println("Db URI >>", dbURI)

	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dbURI)
		if err == nil {
			break
		}

		fmt.Println("Cannot connect to DB, trying again on 5 secs...")
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return nil, err
	}

	return db, nil
}
