package main

import (
	"log"
	"myapp/config"
	"net/http"
	"os"
	"time"

	"myapp/controller"

	"myapp/middleware"
	"myapp/router"

	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client, err := config.NewEntClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	if err != nil {
		log.Println("Fail to initialize client")
	}
	controller.EntClient = client

	r := mux.NewRouter()
	r.Use(middleware.Header)

	router.Register(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on port " + port)
	log.Fatal(srv.ListenAndServe())
}
