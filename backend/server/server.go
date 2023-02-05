package server

import (
	"encoding/json"
	"net/http"

	"github.com/choiivan13/ustchart/backend/internal/db"
	"github.com/choiivan13/ustchart/backend/internal/scraper"
	"github.com/choiivan13/ustchart/backend/types"
)

type Server struct {
	MongoDB db.Operations
	Scraper scraper.Operations
}

func NewServer(offering string, interval float64) *Server {
	mongoDB := db.NewDBHandler()
	scraper := scraper.NewScraper(offering, interval, mongoDB)
	return &Server{
		MongoDB: mongoDB,
		Scraper: scraper,
	}
}

func (s Server) QuerySingleSection(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	query := req.URL.Query()
	offering := query.Get("offering")
	courseName := query.Get("coursename")
	sectionName := query.Get("sectionname")

	if offering == "" || courseName == "" || sectionName == "" {
		return
	}

	data := s.MongoDB.GetSection(&types.Section{
		Offering:    offering,
		CourseName:  courseName,
		SectionName: sectionName,
	})

	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	res.Write(b)
}

func (s Server) GetAllSectionIdentifier(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	data := s.MongoDB.GetSectionsIdentifier()

	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	res.Write(b)
}
