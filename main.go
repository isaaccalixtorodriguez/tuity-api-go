package main

import (
	"log"

	"github.com/isaaccalixtorodriguez/tuity-api-go/database"
	"github.com/isaaccalixtorodriguez/tuity-api-go/env"
	"github.com/isaaccalixtorodriguez/tuity-api-go/router"
)

func main() {
	env.Set()
	if !database.Connection() {
		log.Fatal("Error main database")
		return
	}
	router.SetUp()
}
