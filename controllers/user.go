package controllers

import(
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"log"

	"github.com/gin-gonic/gin"

	"../models"
)

func UserShowByID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	// convert the id type from string to int
    id, err := strconv.Atoi(c.Param("id"))

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

	var users []models.Users
  
	if err := models.MPosGORM.Raw("SELECT id, username from users where id = ?", int64(id)).Scan(&users).Error; err != nil {
		fmt.Printf("error show order by id: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	if (users != nil) {
	  c.JSON(http.StatusOK, users)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	
}

func UserEdit(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var user models.Users
	c.BindJSON(&user)
  
	if err := models.MPosGORM.Model(&user).Where("id = ?", user.Id).Updates(models.Users{Username: user.Username}).Error; err != nil {
		fmt.Printf("error update user : %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	  
	c.JSON(http.StatusOK, gin.H{
	  "user": user.Id,
	  "message": "user edit success",
	})
}