package main

import (
	"fmt"
	"time"

	"github.com/choiivan13/ustchart/backend/internal/db"
	"github.com/choiivan13/ustchart/backend/internal/scraper"
	"github.com/robfig/cron/v3"
)

func main() {
	// s := os.Getenv("USTCHARTURL")
	// fmt.Println(s)

	c := cron.New()
	mongodb := db.NewDBHandler()
	s := scraper.NewScraper("2230", 123, mongodb)

	// section := types.Section{
	// 	Offering:    "2230",
	// 	CourseName:  "BIBU 4830 - Biotechnology Management (3 units)",
	// 	SectionName: "L1 (1730)",
	// }
	// result := s.DB.GetSection(&section)
	// fmt.Println(result)

	c.AddFunc("*/10 * * * *", func() {
		t := time.Now().Truncate(5 * time.Minute)
		fmt.Println("Time now is: ", t.String(), ", scraping...")
		s.Scrape(t.Unix())
		fmt.Println("Job Finished!")
	})
	c.Start()
	time.Sleep(8 * (time.Hour))
}

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
