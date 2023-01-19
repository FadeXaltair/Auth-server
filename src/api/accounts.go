package api

import (
	"net/http"
	"save-the-queen-backend/config"
	"save-the-queen-backend/database"
	"save-the-queen-backend/src/auth"

	"github.com/gin-gonic/gin"
)

// signup function is used to create the account in the database
func SignUp(c *gin.Context) {
	var userdata database.User
	err := c.Bind(&userdata)
	if err != nil {
		config.Error(err)
		return
	}
	password := userdata.Password
	hash, _ := auth.HashPassword(password)
	status := auth.CheckPasswordHash(password, hash)
	if !status {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "incorrect pass",
		})
		return
	}
	data := database.User{
		Name:     userdata.Name,
		Role:     userdata.Role,
		Email:    userdata.Email,
		Password: hash,
	}
	result := config.DB.Create(&data)
	if result.Error != nil {
		config.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "account created successfully",
	})

}

// Login function is used to login the user
func Login(c *gin.Context) {
	var data database.Login
	err := c.Bind(&data)
	if err != nil {
		config.Error(err)
		return
	}

	var user database.User
	res := config.DB.First(&user, "email= ?", data.Email)
	if res.Error != nil {
		config.Error(err)
		return
	}
	status := auth.CheckPasswordHash(data.Password, user.Password)
	if !status {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "incorrect password",
		})
		return
	}
	tokenstring, _ := GetData(data.Email)
	token, _ := auth.GenerateJwtToken(tokenstring.Name, tokenstring.Role)
	resp := database.Response{
		Id:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  resp,
		"token": token,
	})
}

type User struct {
	Role string `json:"role"`
	Name string `json:"name"`
}

func GetData(email string) (User, error) {
	var data User
	err := config.DB.Raw(`select x.role,x.name from public.users x  where email  = '` + email + `'`).Scan(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}
