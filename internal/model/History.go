package model

type Route struct {
	ID             int     `json:"id,omitempty"`
	LongitudeStart float64 `json:"longitudeStart" gorm:"longitudeStart"`
	LatitudeStart  float64 `json:"latitudeStart" gorm:"latitudeStart"`
	LongitudeEnd   float64 `json:"longitudeEnd" gorm:"longitudeEnd"`
	LatitudeEnd    float64 `json:"latitudeEnd" gorm:"latitudeEnd"`
	BranchID       int     `json:"branchId,omitempty" gorm:"branchID"`
	AtmID          int     `json:"atmId,omitempty" gorm:"atmId"`
	UserID         int     `json:"userId" gorm:"userId"`
}
