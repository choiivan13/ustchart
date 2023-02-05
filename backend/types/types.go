package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// DB
type Section struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Offering    string             `json:"offering,omitempty" bson:"offering,omitempty"`
	CourseName  string             `json:"coursename,omitempty" bson:"coursename,omitempty"`
	SectionName string             `json:"sectionname,omitempty" bson:"sectionname,omitempty"`
	CourseTime  string             `json:"coursetime,omitempty" bson:"coursetime,omitempty"`
	Instructors []string           `json:"instructors,omitempty" bson:"instructors,omitempty"`
	Data        Data               `json:"data,omitempty" bson:"data,omitempty"`
}

type Data struct {
	TimeStamp int64 `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Quota     int   `json:"quota,omitempty" bson:"quota,omitempty"`
	Enrol     int   `json:"enrol,omitempty" bson:"enrol,omitempty"`
	Wait      int   `json:"wait,omitempty" bson:"wait,omitempty"`
}

// Frontend
type SectionData struct {
	Offering    string   `json:"offering,omitempty" bson:"offering,omitempty"`
	CourseName  string   `json:"coursename,omitempty" bson:"coursename,omitempty"`
	SectionName string   `json:"sectionname,omitempty" bson:"sectionname,omitempty"`
	CourseTime  string   `json:"coursetime,omitempty" bson:"coursetime,omitempty"`
	Instructors []string `json:"instructors,omitempty" bson:"instructors,omitempty"`
	Data        []Data   `json:"data,omitempty" bson:"data,omitempty"`
}

type SectionIdentifier struct {
	Offering    string `json:"offering,omitempty" bson:"offering,omitempty"`
	CourseName  string `json:"coursename,omitempty" bson:"coursename,omitempty"`
	SectionName string `json:"sectionname,omitempty" bson:"sectionname,omitempty"`
}
