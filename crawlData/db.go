package crawlData

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Caoduc123@tcp(127.0.0.1:3306)/crawler")

	if err != nil {
		fmt.Println(err)
	}

	return db, nil
}
