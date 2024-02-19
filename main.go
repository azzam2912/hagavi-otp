package main

import (
	"hagavi-otp/database"
	"hagavi-otp/router"
	"log"
)

func main() {
	db := database.Connect()
	app := router.New(db)
	log.Fatal(app.Listen(":3000"))
}