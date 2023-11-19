package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv)
	}
	mysqlUsername := os.Getenv("MYSQLUSERNAME")
	mysqlPort := os.Getenv("DATABASEPORT")
	var err error
	const MYSQL = "%s:@tcp(127.0.0.1:%s)/facetrack?charset=utf8mb4&parseTime=True&loc=Local"
	connectionString := fmt.Sprintf(MYSQL, mysqlUsername, mysqlPort)
	dsn := connectionString
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("CANT CONNECT TO DATABASE")
	}

	fmt.Println("YOU'RE CONNECTED TO DATABASE")

}