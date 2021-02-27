package controllers

import(
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"../configs"
	"../models"
	"../utils"
)

func Process(c *gin.Context){
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	typ := "update"
	var user models.Users
  
	if err := configs.MPosGORM.Model(&user).Where("type = ?", typ).Updates(models.Users{Username: utils.RandomString(10)}).Error; err != nil {
		fmt.Printf("process error : %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	  
	c.JSON(http.StatusOK, gin.H{
	  "message": "process success",
	})
}