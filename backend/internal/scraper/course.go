package scraper

type Department struct {
	DepartmentName string
	Courses        []Course
}

type Course struct {
	CourseName string
	Sections   []Section
}

type Section struct {
	SectionName string
	Time        string
	Instructors []string
	Quota       int
	Enrol       int
	Wait        int
}
