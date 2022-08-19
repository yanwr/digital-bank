package main

import (
	"yanwr/digital-bank/config"
	"yanwr/digital-bank/database"
)

func init() {
	config.LoadEnv()
}

func main() {
	database.LoadConnectionDB()

	server := config.CreateServer()
	server.LoadRoutesAndRunServer()
}
