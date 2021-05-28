package main

import (
	"github.com/dileofrancoj/blog-app/config"
	"github.com/dileofrancoj/blog-app/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	config.ConnectDatabase()
	routes.Routes()

}