package main

import (
	"fmt"
	c "go-module/crawlData"
	"log"
	// "go-module/database"
)

func main() {
	// c.Crawler()
	db, err := c.ConnectDatabase()

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	log.Println("Connect Success")
	c.Crawler(db)

}
