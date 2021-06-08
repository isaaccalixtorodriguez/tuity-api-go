package database

import (
	"context"
	"fmt"
	"log"

	"github.com/isaaccalixtorodriguez/tuity-api-go/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() bool {
	url := fmt.Sprintf("mongodb+srv://%v:%v@cluster0.kxixs.mongodb.net/%v?retryWrites=true&w=majority", env.Get("DB_USER", "user"), env.Get("DB_PASSWORD", "pass"), env.Get("DB_NAME", "db"))
	optionsClient := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), optionsClient)
	if err != nil {
		log.Fatalf("Error in the connection with mongodb : %v", err.Error())
		return false
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(url)
		log.Fatalf("Error in the ping with mongodb : %v", err.Error())
		return false
	}

	log.Println("Successful connection to mongodb")
	return true
}
