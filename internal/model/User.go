package model

type User struct {
	User_ID        int    `json:"-" gorm:"primarykey;autoIncrement"`
	FirstName      string `json:"FirstName"`
	SecondName     string `json:"SecondName"`
	MiddleName     string `json:"MiddleName,omitempty"`
	Email          string `json:"Email"`
	Password       string `json:"Password" gorm:"column:Password"`
	RepeatPassword string `json:"rPassword" gorm:"-"`
	IsConfirmed    bool   `json:"-"`
	AccessToken    string `json:"-" gorm:"column:accesstoken"`
	LegalEntity    bool   `json:"LegalEntity" gorm:"column:legalentity"`
}
