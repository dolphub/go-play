package main

import (
	"github.com/dolphub/api-server/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server.StartServer()
}
