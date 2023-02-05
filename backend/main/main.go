package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/choiivan13/ustchart/backend/server"
	"github.com/robfig/cron/v3"
)

func main() {
	// mongodb := db.NewDBHandler()
	// s := scraper.NewScraper("2230", 123, mongodb)
	s := server.NewServer("2230", 123)

	// Scraper cron
	go func() {
		c := cron.New()
		c.AddFunc("*/10 * * * *", func() {
			t := time.Now().Truncate(5 * time.Minute)
			fmt.Println("Time now is: ", t.String(), ", scraping...")
			s.Scraper.Scrape(t.Unix())
			fmt.Println("Job Finished!")
		})
		c.Start()
		for time.Now().Before(time.Date(2023, 2, 15, 0, 0, 0, 0, time.Now().Location())) {
			time.Sleep(8 * (time.Hour))
		}
	}()

	// Http Listening
	http.HandleFunc("/section", s.QuerySingleSection)
	http.HandleFunc("/sections", s.GetAllSectionIdentifier)

	log.Fatal(http.ListenAndServe(":3000", nil))
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
