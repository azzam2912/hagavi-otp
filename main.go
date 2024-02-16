package main

import (
	"fmt"
	"hagavi-otp/database"
	"hagavi-otp/router"
	"log"
)

func main() {
	app := router.New()
	err := database.Connect()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(app.Listen(":3000"))
}