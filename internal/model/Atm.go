package model

type Atms struct {
	ID                                         int     `gorm:"column:id;primary_key"`
	Address                                    string  `gorm:"column:address"`
	Latitude                                   float64 `gorm:"column:latitude"`
	Longitude                                  float64 `gorm:"column:longitude"`
	AllDay                                     string  `gorm:"column:allDay"`
	ServicesWheelchairServiceCapability        string  `gorm:"column:serviceswheelchairservicecapability"`
	ServicesWheelchairServiceActivity          string  `gorm:"column:serviceswheelchairserviceactivity" json:"serviceswheelchairserviceActivity"`
	ServicesBlindServiceCapability             string  `gorm:"column:servicesblindservicecapability"`
	ServicesBlindServiceActivity               string  `gorm:"column:servicesblindserviceactivity"`
	ServicesNFCForBankCardsServiceCapability   string  `gorm:"column:servicesnfcforbankcardsservicecapability"`
	ServicesNFCForBankCardsServiceActivity     string  `gorm:"column:servicesnfcforbankcardsserviceactivity"`
	ServicesQRReadServiceCapability            string  `gorm:"column:servicesqrreadservicecapability"`
	ServicesQRReadServiceActivity              string  `gorm:"column:servicesqrreadserviceactivity"`
	ServicesSupportsUSDServiceCapability       string  `gorm:"column:servicessupportsusdservicecapability"`
	ServicesSupportsUSDServiceActivity         string  `gorm:"column:servicessupportsusdserviceactivity"`
	ServicesSupportsChargeRUBServiceCapability string  `gorm:"column:servicessupportschargerubservicecapability"`
	ServicesSupportsChargeRUBServiceActivity   string  `gorm:"column:servicessupportschargerubserviceactivity"`
	ServicesSupportsEURServiceCapability       string  `gorm:"column:servicessupportseurservicecapability"`
	ServicesSupportsEURServiceActivity         string  `gorm:"column:servicessupportseurserviceactivity"`
	ServicesSupportsRUBServiceCapability       string  `gorm:"column:servicessupportsrubservicecapability"`
	ServicesSupportsRUBServiceActivity         string  `gorm:"column:servicessupportsrubserviceactivity"`
}

type ATMResponse struct {
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	AllDay    bool    `json:"allDay"`
	Services  struct {
		Wheelchair struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"wheelchair"`
		Blind struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"blind"`
		NFCForBankCards struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"nfcForBankCards"`
		QRRead struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"qrRead"`
		SupportsUSD struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"supportsUsd"`
		SupportsChargeRUB struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"supportsChargeRub"`
		SupportsEUR struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"supportsEur"`
		SupportsRUB struct {
			ServiceCapability string `json:"serviceCapability"`
			ServiceActivity   string `json:"serviceActivity"`
		} `json:"supportsRub"`
	} `json:"services"`
}
