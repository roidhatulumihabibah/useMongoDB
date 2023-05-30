package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func main() {
	// Konfigurasi koneksi MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Membaca data dari koleksi
	collection := client.Database("dosen").Collection("person")
	cur, err := collection.Find(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	// Menampilkan data
	for cur.Next(context.Background()) {
		var person Person
		err := cur.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name: %s, Email: %s\n", person.Name, person.Email)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
