package main

import (
	"github.com/Tambarie/payment-gateway/application/server"
	"github.com/Tambarie/payment-gateway/domain/helpers"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	helpers.InitializeLogDir()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env with godotenv: %s", err)
	}
	server.Start()
}
