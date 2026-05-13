package main

import (
	"log"

	httpserver "iteraciones/backend/internal/http"
)

func main() {
	router := httpserver.NewRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
