package model

type User struct {
	User_ID        int    `json:"User_ID" gorm:"primarykey;autoIncrement"`
	FirstName      string `json:"FirstName"`
	SecondName     string `json:"SecondName"`
	MiddleName     string `json:"MiddleName"`
	Email          string `json:"Email"`
	Password       string `json:"Password" gorm:"column:Password"`
	RepeatPassword string `json:"rPassword" gorm:"-"`
	IsConfirmed    bool   `json:"confirmationCode"`
	AccessToken    string `gorm:"column:accesstoken"`
}

type Talon struct {
	ID          int  `json:"id,omitempty"`
	LegalEntity bool `json:"legalEntity"`
	UserID      int  `json:"userId"`
	BranchID    int  `json:"branchId"`
}
