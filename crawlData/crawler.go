package crawlData

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Data struct {
	Name string
	Year string
	Rate string
	Link string
}

func Crawler(db *sql.DB) {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visitting : %s \n", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		checkErr(e)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Visited : %s \n", r.Request.URL)
	})

	c.OnHTML("tr", func(h *colly.HTMLElement) {
		data := Data{}
		data.Name = h.ChildText(".titleColumn a")
		data.Year = h.ChildText(".titleColumn span")
		data.Rate = h.ChildText(".ratingColumn strong")
		data.Link = h.ChildAttr("a", "href")
		fmt.Printf("Name : %s\n Year : %s\n Rate : %s\n Link : %s\n", data.Name, data.Year, data.Rate, data.Link)

		//insert to db
		st, err := db.Prepare("insert into data (name,year,rate,link) values(?,?,?,?)")
		checkErr(err)
		res, _ := st.Exec(data.Name, data.Year, data.Rate, data.Link)

		// _, err1 := res.LastInsertId()
		// checkErr(err1)
		_, err2 := res.LastInsertId()
		if err != nil {
			log.Fatal(err2)
		}

		// fmt.Printf("=&gt;Insert ID: %d", lastId) //In ra màn hình ID vừa insert
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")
}
