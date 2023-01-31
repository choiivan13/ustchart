package main

import (
	"time"

	"github.com/choiivan13/ustchart/backend/internal/db"
	"github.com/choiivan13/ustchart/backend/internal/scraper"
)

func main() {
	// s := os.Getenv("USTCHARTURL")
	// fmt.Println(s)
	mongodb := db.NewDBHandler()
	c := scraper.NewScraper("2230", 123, mongodb)
	c.Scrape(time.Now().Unix())

	// s := types.Section{
	// 	Offering:    "2230",
	// 	CourseName:  "TEST 1000",
	// 	SectionName: "L1 (1001)",
	// 	CourseTime:  "Fri 10:00",
	// 	Instructors: []string{
	// 		"Ivan",
	// 		"Max",
	// 	},
	// 	Data: types.Data{
	// 		TimeStamp: time.Now().Unix(),
	// 		Quota:     100,
	// 		Enrol:     20,
	// 		Wait:      30,
	// 	},
	// }
	// mongodb.UpdateSection(s)
}
