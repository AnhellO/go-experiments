package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Post represents a blog post
type Post struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Body      string             `json:"body,omitempty" bson:"body,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

var collection *mongo.Collection

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")
	collection = client.Database("test_db").Collection("posts")
}

func main() {
	// Insert one post
	InsertPost("Pruebas!", "<p>Probando ando!</p>")
	// Fetch one post
	post := GetPost()
	fmt.Println(post)
}

// InsertPost will insert a new post record into the mongo collection
func InsertPost(title string, body string) {
	post := Post{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     title,
		Body:      body,
	}
	insertResult, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted post with ID:", insertResult.InsertedID)
}

// GetPost will return a single record from the mongo collection
func GetPost() Post {
	filter := bson.D{}

	var post Post
	err := collection.FindOne(context.TODO(), filter).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found post with title ", post.Title)
	return post
}
