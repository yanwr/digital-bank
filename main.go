package main

import (
	"yanwr/digital-bank/database"
	"yanwr/digital-bank/env"
	server "yanwr/digital-bank/serve"
)

const (
	ENV             = "ENV"
	DB_HOST         = "DB_HOST"
	DB_USER         = "DB_USER"
	DB_PASS         = "DB_PASS"
	DB_NAME         = "DB_NAME"
	DB_TZ           = "DB_TZ"
	DB_TYPE         = "DB_TYPE"
	DB_AUTO_MIGRATE = "DB_AUTO_MIGRATE"

	SERVER_PORT = "SERVER_PORT"
)

func init() {
	env.LoadEnv()
}

func main() {
	database.LoadConnectionDB()

	server := server.CreateServer(database.GetConnectionDB())
	server.RunServerAndLoadRoutes()
}
