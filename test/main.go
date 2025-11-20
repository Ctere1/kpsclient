package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Ctere1/kpsclient"
	"github.com/joho/godotenv"
)

// Example main for querying the KPS service, with query parameters loaded from .env.
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}
	username := os.Getenv("KPS_USERNAME")
	password := os.Getenv("KPS_PASSWORD")
	if username == "" || password == "" {
		log.Fatal("Please set KPS_USERNAME and KPS_PASSWORD environment variables")
	}

	// Load query parameters from environment variables (or use dummy defaults if empty)
	req := kpsclient.QueryRequest{
		TCNo:         os.Getenv("KPS_TCNO"),
		FirstName:    os.Getenv("KPS_FIRSTNAME"),
		LastName:     os.Getenv("KPS_LASTNAME"),
		BirthYear:    os.Getenv("KPS_BIRTHYEAR"),
		BirthMonth:   os.Getenv("KPS_BIRTHMONTH"),
		BirthDay:     os.Getenv("KPS_BIRTHDAY"),
		SerialNumber: os.Getenv("KPS_SERIALNUMBER"),
	}

	client := kpsclient.New(username, password, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	res, err := client.DoQuery(ctx, req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	PrintKPSResultBox(res)
}
