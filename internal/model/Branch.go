package model

type BusinessResponse struct {
	ID         int `json:"id"`
	TalonCount int `json:"talonCount"`
}

type Branch struct {
	Branch_ID           int     `gorm:"column:id;primary_key"`
	SalePointName       string  `gorm:"column:salePointName"`
	Address             string  `gorm:"column:address"`
	Status              string  `gorm:"column:status"`
	OpenHours           []byte  `gorm:"column:openHours" sql:"type:json"`
	RKO                 string  `gorm:"column:rko"`
	OpenHoursIndividual []byte  `gorm:"column:openHoursIndividual" sql:"type:json"`
	OfficeType          string  `gorm:"column:officeType"`
	SalePointFormat     string  `gorm:"column:salePointFormat"`
	SUOAvailability     string  `gorm:"column:suoAvailability"`
	HasRamp             string  `gorm:"column:hasRamp"`
	Latitude            float64 `gorm:"column:latitude"`
	Longitude           float64 `gorm:"column:longitude"`
	MetroStation        string  `gorm:"column:metroStation"`
	Distance            int64   `gorm:"column:distance"`
	KEP                 *bool   `gorm:"column:kep"`
	MyBranch            bool    `gorm:"column:myBranch"`
	Network             string  `gorm:"column:network"`
	SalePointCode       string  `gorm:"column:salePointCode"`
	TalonCount          int     `gorm:"column:talonCount"`
}

type BranchResponse struct {
	Branch_ID           int             `json:"id"`
	SalePointName       string          `json:"salePointName"`
	Address             string          `json:"address"`
	Status              string          `json:"status"`
	OpenHours           []OpenHoursType `json:"openHours"`
	RKO                 string          `json:"rko"`
	OpenHoursIndividual []OpenHoursType `json:"openHoursIndividual"`
	OfficeType          string          `json:"officeType"`
	SalePointFormat     string          `json:"salePointFormat"`
	SUOAvailability     string          `json:"suoAvailability"`
	HasRamp             string          `json:"hasRamp"`
	Latitude            float64         `json:"latitude"`
	Longitude           float64         `json:"longitude"`
	MetroStation        string          `json:"metroStation"`
	Distance            int64           `json:"distance"`
	KEP                 *bool           `json:"kep"`
	MyBranch            bool            `json:"myBranch"`
	Network             string          `json:"network"`
	SalePointCode       string          `json:"salePointCode"`
	TalonCount          int             `json:"talonCount"`
}

type OpenHoursType struct {
	Days  string `json:"days"`
	Hours string `json:"hours"`
}
