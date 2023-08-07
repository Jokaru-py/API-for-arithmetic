package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Jokaru-py/API-for-arithmetic/handle"
	"github.com/Jokaru-py/API-for-arithmetic/internal/app"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.UserAccessAuthentication)

	router.HandleFunc("/api/operation/", handle.Operation)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, router) //Запустите приложение, посетите localhost:8000/api

	if err != nil {
		fmt.Print(err)
	}
}
