package main

import (
	"time"
	"fmt"
	"os"
  
	"github.com/gin-gonic/gin"
  
	"../simple-app-golang/controllers"
	"../simple-app-golang/models"

)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	fmt.Printf("Started at : %3v \n", time.Now())
  
	//InitPostgres()
	models.InitGormPostgres()
	defer models.MPosGORM.Close()
  
	// Set the router as the default one shipped with Gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
  
	// Setup route group for the API
	api := router.Group("/api")
  
	api.POST("/users/edit", controllers.UserEdit)
	api.GET("/users/id/:id", controllers.UserShowByID) 
  
	// Start and run the server
	fmt.Printf("App listening at : 4000")
	router.Run(":4000")
  }