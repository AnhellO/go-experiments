package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Post represents a blog post
type Post struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `json:"title,omitempty" bson:"title,omitempty"`
	Body       string             `json:"body,omitempty" bson:"body,omitempty"`
	User       User               `json:"user,omitempty" bson:"user,omitempty"`
	Taxonomies []Taxonomy         `json:"taxonomies,omitempty" bson:"taxonomies,omitempty"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

// User represents a user who wrotes the blog post
type User struct {
	FirstName string    `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Level     int       `json:"level,omitempty" bson:"level,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// Taxonomy represents a blog post taxonomy
type Taxonomy struct {
	Type      string    `json:"type,omitempty" bson:"type,omitempty"`
	Values    []string  `json:"values,omitempty" bson:"values,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
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
	gofakeit.Seed(0)
	// Insert some randomly generated posts
	for i := 0; i < 10; i++ {
		InsertPost(User{})
	}

	// Intentionally insert a new user that we will look for
	InsertPost(User{FirstName: "Chano", LastName: "Menguano", Level: 5})
	// Fetch one post
	post := GetPost(bson.M{"user.first_name": "Chano"}, bson.M{"title": 1, "user.first_name": 1, "user.last_name": 1})
	fmt.Println(post)
}

// InsertPost will insert a new post record into the mongo collection
func InsertPost(user User) {
	post := generatePost(user)
	insertResult, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted post with ID:", insertResult.InsertedID)
}

// GetPost will return a single record from the mongo collection
func GetPost(filter, projection bson.M) Post {
	var post Post
	err := collection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found post with title ", post.Title)
	return post
}

func generatePost(user User) Post {
	post := Post{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     gofakeit.Sentence(10),
		Body:      gofakeit.LoremIpsumParagraph(2, 5, 50, " "),
		User: User{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Level:     gofakeit.Number(1, 5),
		},
		Taxonomies: []Taxonomy{
			{
				Type:   gofakeit.Word(),
				Values: []string{gofakeit.Word(), gofakeit.Word()},
			},
			{
				Type:   gofakeit.Word(),
				Values: []string{gofakeit.Word(), gofakeit.Word()},
			},
		},
	}

	if user != (User{}) {
		post.User = user
	}

	return post
}
