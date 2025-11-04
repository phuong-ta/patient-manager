package main

import (
	"log"

	"patient-manager/internal/app"
)

func main() {
	srv := app.NewServer()
	if err := srv.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
