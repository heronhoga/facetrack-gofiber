package migration

import (
	"facetrack-gofiber/database"
	"facetrack-gofiber/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Prediction{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database migrated")
}