package main

import (
	"log"

	"github.com/joho/godotenv"
)

func initEnv() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	initEnv()
}
