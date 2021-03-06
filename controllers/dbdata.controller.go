package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"../configs"
	"../models"
	"../utils"
)

func Process(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	key := c.Param("key")

	var dbDatas []models.DbData

	if err := configs.MPosGORM.Raw("SELECT * from db_data where key = ?", key).Scan(&dbDatas).Error; err != nil {
		fmt.Printf("error show dbData by key: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if dbDatas != nil {
		for _, dData := range dbDatas {
			if err := configs.MPosGORM.Model(&dData).Where("key = ?", dData.Key).Updates(models.DbData{Content: utils.RandomString(10), UpdatedOn: time.Now()}).Error; err != nil {
				fmt.Printf("Process Error : %3v \n", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"key":     key,
			"message": "Process Success",
		})
	} else {
		c.JSON(http.StatusOK, json.RawMessage(`Process Fail`))
	}
}
