package main

import (
	"facetrack-gofiber/database"
	"facetrack-gofiber/route"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
    //Connecting to database
	database.DatabaseInit()

    //Running migration
    // migration.RunMigration()

    //Using godotenv
    errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv)
	}

    port := os.Getenv("APP_PORT")
    app := fiber.New()
    route.RouteInit(app)
    app.Listen(port)

}
