package main

import (
	"fmt"
	"log"
  
	"../simple-app-golang/configs"
	"../simple-app-golang/routers"

)

func main() {
	//InitPostgres()
	configs.InitGormPostgres()
	defer configs.MPosGORM.Close()
	router := routers.InitRouter()
  
	// Start and run the server
	port := "4000"
	fmt.Printf("App listening on port "+port)
	log.Fatal(router.Run(":"+port))
  }