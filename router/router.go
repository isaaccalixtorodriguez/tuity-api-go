package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/isaaccalixtorodriguez/tuity-api-go/env"
	"github.com/rs/cors"
)

func SetUp() {
	router := mux.NewRouter()

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+env.Get("PORT", "3000"), handler))
}
