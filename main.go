package main

import (
	"yanwr/digital-bank/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	println("Digital Bank works !!")
	config.ConnectDB()

	// config.ConnectDB().Create(&models.Account{})
}
