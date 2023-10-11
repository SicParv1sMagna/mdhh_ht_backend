package model

import "github.com/google/uuid"

type User struct {
	User_ID        uuid.UUID `json:"User_ID" gorm:"primarykey;autoIncrement"`
	FirstName      string    `json:"FirstName"`
	SecondName     string    `json:"SecondName"`
	MiddleName     string    `json:"MiddleName"`
	Email          string    `json:"Email" gorm:"column:Login"`
	Password       string    `json:"Password" gorm:"column:Password"`
	RepeatPassword string    `json:"rPassword" gorm:"-"`
}
