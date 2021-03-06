package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"../simple-app-golang/configs"
	"../simple-app-golang/routers"
)

func main() {
	//Load Environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//InitPostgres()
	configs.InitGormPostgres()
	defer configs.MPosGORM.Close()
	router := routers.InitRouter()

	// Start and run the server
	port := "4000"
	fmt.Printf("App listening on port " + port)
	log.Fatal(router.Run(":" + port))
}
