package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/choiivan13/ustchart/backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Operations interface {
	UpdateSection(s *types.Section)
	GetSection(s *types.Section) types.SectionData
	GetSectionsIdentifier() []types.SectionIdentifier
}

type MongoDBHandler struct {
	Client *mongo.Client
}

// !!!
func NewDBHandler() *MongoDBHandler {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("USTCHARTURL")).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	return &MongoDBHandler{
		Client: client,
	}
}

func (m MongoDBHandler) UpdateSection(s *types.Section) {
	collection := m.Client.Database("ustchart").Collection(s.Offering)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{
		"offering":    s.Offering,
		"coursename":  s.CourseName,
		"sectionname": s.SectionName,
	}
	update := bson.M{
		"$push": bson.M{"data": s.Data},
		"$setOnInsert": bson.M{
			"coursetime":  s.CourseTime,
			"instructors": s.Instructors,
		},
	}
	upsert := true
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := collection.UpdateOne(ctx, filter, update, &opt)
	if err != nil {
		panic(err)
	}
}

func (m MongoDBHandler) GetSection(s *types.Section) types.SectionData {
	collection := m.Client.Database("ustchart").Collection(s.Offering)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{
		"offering":    s.Offering,
		"coursename":  s.CourseName,
		"sectionname": s.SectionName,
	}
	options := options.FindOne().SetProjection(bson.M{"_id": 0})
	var result types.SectionData
	err := collection.FindOne(ctx, filter, options).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func (m MongoDBHandler) GetSectionsIdentifier() []types.SectionIdentifier {
	collection := m.Client.Database("ustchart").Collection("2230")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Find().SetProjection(bson.M{"offering": 1, "coursename": 1, "sectionname": 1, "_id": 0})

	cur, err := collection.Find(ctx, bson.D{}, options)
	if err != nil {
		panic(err)
	}

	var results []types.SectionIdentifier
	cur.All(context.Background(), &results)
	// for cur.Next(context.Background()) {
	// 	// To decode into a struct, use cursor.Decode()
	// 	err := cur.Decode(&result)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	return results
}
