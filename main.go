package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rudiarta/refactory_test/api/route"
	"github.com/rudiarta/refactory_test/container"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	container := container.BuildContainer()
	if er := container.Invoke(route.InvokeRoute); er != nil {
		panic(er)
	}
}
