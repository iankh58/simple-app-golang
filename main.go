package main

import (
	"time"
	"fmt"
	"os"
  
	"../simple-app-golang/configs"
	"../simple-app-golang/routers"

)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	fmt.Printf("Started at : %3v \n", time.Now())
  
	//InitPostgres()
	configs.InitGormPostgres()
	defer configs.MPosGORM.Close()
	router := routers.InitRouter()
  
	// Start and run the server
	fmt.Printf("App listening at : 4000\n")
	router.Run(":4000")
  }