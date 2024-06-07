package main

import (
	"log"

	"github.com/Nier704/url-shortener/internal/db"
	"github.com/Nier704/url-shortener/internal/handlers"
	routes "github.com/Nier704/url-shortener/internal/routers"
)

func main() {
	db, err := db.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handler := handlers.NewUrlHandler(db)
	router := routes.NewRouter(handler)

	router.Init()
	router.Start()
}
