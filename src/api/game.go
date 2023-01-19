package api

import (
	"log"
	"net/http"
	"save-the-queen-backend/config"
	"save-the-queen-backend/database"

	"github.com/gin-gonic/gin"
)

func GameCreate(c *gin.Context) {
	var data database.Game
	err := c.Bind(&data)
	if err != nil {
		log.Println(err)
	}
	role := c.Keys["Role"]
	if role == "admin" {
		err := CreateGame(data)
		if err != nil {
			log.Println(err)
			c.JSON(400, gin.H{
				"error":   true,
				"message": "error while adding data",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"error":   false,
			"message": "game created successfully",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"error" : true,
			"message": "You are not authorised",
		})
	}
}

func ListGame(c *gin.Context) {
	data, err := ListGames()
	if err != nil {
		c.JSON(400, gin.H{
			"error":   false,
			"message": "error while fetching values from database",
		})
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateGame(data database.Game) error {
	result := config.DB.Create(&data)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func ListGames() ([]database.Game, error) {
	var data []database.Game
	err := config.DB.Raw(`select * from public.games`).Scan(&data).Error
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}
