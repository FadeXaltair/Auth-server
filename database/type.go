package database

import (
	"gorm.io/gorm"
)

// Users struct is used to store the data from the user to database
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login struct is used to login in the system
type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Game struct {
	gorm.Model
	AdminId         uint   `json:"admin_id" binding:"required"`
	Client          string `json:"client" binding:"required"`
	ContactPerson   string `json:"contact_person" binding:"required"`
	Partner         string `json:"partner"`
	Date            string `json:"date" binding:"required"`
	TimeZone        string `json:"time_zone" binding:"required"`
	Time            string `json:"time" binding:"required"`
	Teams           int    `json:"teams" binding:"required"`
	TotalHints      int    `json:"total_hints" binding:"required"`
	HintTimePenalty string `json:"hint_time_penalty" binding:"required"`
	GameDuration    string `json:"game_duration" binding:"required"`
}
// Response struct is used for the response
type Response struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Tokens struct {
	Token string `json:"token"`
}
