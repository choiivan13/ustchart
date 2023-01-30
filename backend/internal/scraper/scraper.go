package scraper

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly/v2"
)

type Operations interface {
}

type Scraper struct {
	// Offering will be in the format of YYXX, where YY denotes the first year (22-23), and 10, 20, 30, 40 (Fall, Winter, Spring, Summer)
	Offering string
	// Interval of Scraping Time, measured in minutes
	Interval float64
	Courses  []Course
}

func NewScraper(offering string, interval float64) *Scraper {
	return &Scraper{
		Offering: offering,
		Interval: interval,
		Courses:  make([]Course, 0),
	}
}

func (s Scraper) Scrape() {
	deptCollector := colly.NewCollector(
		colly.Async(),
	)
	courseCollector := deptCollector.Clone()

	deptCollector.OnHTML(".depts .ug", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		// fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		courseCollector.Visit(e.Request.AbsoluteURL(link))
	})

	deptCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	courseCollector.OnHTML(".course", func(e *colly.HTMLElement) {
		// fmt.Println("Course found", e.ChildText("h2"))
		courseName := e.ChildText("h2")
		if courseName == "" {
			log.Println("No title found", e.Request.URL)
		}
		course := Course{
			CourseName: courseName,
		}
		// Iterate over rows of the table which contains different information
		// about the course
		e.ForEach(".newsect", func(_ int, newSection *colly.HTMLElement) {
			// Instructor
			var instructors []string
			newSection.ForEach("td:nth-child(4) a", func(i int, h *colly.HTMLElement) {
				instructors = append(instructors, h.Text)
			})
			// Quota
			var quotaString string
			if newSection.ChildText(".quotadetail") != "" {
				quotaString = newSection.ChildText("td:nth-child(5) span")
			} else {
				quotaString = newSection.ChildText("td:nth-child(5)")
			}
			quota, err := strconv.Atoi(quotaString)
			if err != nil {
				log.Println(err, courseName)
			}

			// Enrol
			enrol, err := strconv.Atoi(newSection.ChildText("td:nth-child(6)"))
			if err != nil {
				log.Println(err, courseName)
			}

			// Wait
			wait, err := strconv.Atoi(newSection.ChildText("td:nth-child(8)"))
			if err != nil {
				log.Println(err, courseName)
			}

			nextSection := Section{
				SectionName: newSection.ChildText("td:nth-child(1)"),
				Time:        newSection.ChildText("td:nth-child(2)"),
				Instructors: instructors,
				Quota:       quota,
				Enrol:       enrol,
				Wait:        wait,
			}
			course.Sections = append(course.Sections, nextSection)

			s.Courses = append(s.Courses, course)
			fmt.Println("Section added:", nextSection)
		})
		fmt.Println("Course finished: ", courseName)
	})

	deptCollector.Visit("https://w5.ab.ust.hk/wcq/cgi-bin/2210/")
	deptCollector.Wait()
	courseCollector.Wait()
}
