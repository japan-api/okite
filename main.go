package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bgentry/heroku-go"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := heroku.Client{Username: os.Getenv("HEROKU_LOGIN"), Password: os.Getenv("HEROKU_TOKEN")}

	resp, err := http.Get("http://japan-api.ninja")
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 503 {
		client.DynoRestartAll("japan-api")
	}

}
