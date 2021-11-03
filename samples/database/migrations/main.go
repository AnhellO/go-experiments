package main

import (
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func init() {
	m, err := migrate.New(
		"file://db/migrations",
		"mongodb://localhost:27017/migrations-test",
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func main() {

}
