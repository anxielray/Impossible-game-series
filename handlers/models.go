package handlers

import (
    "github.com/gofiber/fiber/v2"
)

var App = fiber.New()

type Answer struct {
	First  int `form:"first"`
	Second int `form:"second"`
	Third  int `form:"third"`
	Fourth int `form:"fourth"`
}

type  Answers struct {
	Ans []Answer
}


type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}