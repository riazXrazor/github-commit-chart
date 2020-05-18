package main

//go:generate go run main.go

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// username := "riazXrazor"
	// repo := "udemy-dl"

	// data := GetRepoCommitData(username, repo)

	// GenerateChart(data, username, repo)

	ServerInit()
}
