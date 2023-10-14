package delivery

import (
	"net/http"
	"strconv"

	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
	"github.com/SicParv1sMagna/mdhh_backend/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetAllAtms(repository *repository.Repository, c *gin.Context) {
	var atms []model.Atms

	atms, err := repository.GetAllAtms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var atmResponses []model.ATMResponse
	for _, atm := range atms {
		atmResponse := model.ATMResponse{
			Address:   atm.Address,
			Latitude:  atm.Latitude,
			Longitude: atm.Longitude,
			AllDay:    atm.AllDay == "true",
		}

		atmResponse.Services.Wheelchair.ServiceCapability = atm.ServicesWheelchairServiceCapability
		atmResponse.Services.Wheelchair.ServiceActivity = atm.ServicesWheelchairServiceActivity

		atmResponse.Services.Blind.ServiceCapability = atm.ServicesBlindServiceCapability
		atmResponse.Services.Blind.ServiceActivity = atm.ServicesBlindServiceActivity

		atmResponse.Services.NFCForBankCards.ServiceCapability = atm.ServicesNFCForBankCardsServiceCapability
		atmResponse.Services.NFCForBankCards.ServiceActivity = atm.ServicesNFCForBankCardsServiceActivity

		atmResponse.Services.QRRead.ServiceCapability = atm.ServicesQRReadServiceCapability
		atmResponse.Services.QRRead.ServiceActivity = atm.ServicesQRReadServiceActivity

		atmResponse.Services.SupportsUSD.ServiceCapability = atm.ServicesSupportsUSDServiceCapability
		atmResponse.Services.SupportsUSD.ServiceActivity = atm.ServicesSupportsUSDServiceActivity

		atmResponse.Services.SupportsChargeRUB.ServiceCapability = atm.ServicesSupportsChargeRUBServiceCapability
		atmResponse.Services.SupportsChargeRUB.ServiceActivity = atm.ServicesSupportsChargeRUBServiceActivity

		atmResponse.Services.SupportsEUR.ServiceCapability = atm.ServicesSupportsEURServiceCapability
		atmResponse.Services.SupportsEUR.ServiceActivity = atm.ServicesSupportsEURServiceActivity

		atmResponse.Services.SupportsRUB.ServiceCapability = atm.ServicesSupportsRUBServiceCapability
		atmResponse.Services.SupportsRUB.ServiceActivity = atm.ServicesSupportsRUBServiceActivity

		atmResponses = append(atmResponses, atmResponse)
	}

	c.JSON(http.StatusOK, atmResponses)
}

func GetAtmById(repository *repository.Repository, c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID не может быть отрицательным",
		})
		return
	}

	atm, err := repository.GetAtmById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	atmResponse := model.ATMResponse{
		Address:   atm.Address,
		Latitude:  atm.Latitude,
		Longitude: atm.Longitude,
		AllDay:    atm.AllDay == "true",
	}

	// Маппинг данных из Services
	atmResponse.Services.Wheelchair.ServiceCapability = atm.ServicesWheelchairServiceCapability
	atmResponse.Services.Wheelchair.ServiceActivity = atm.ServicesWheelchairServiceActivity

	atmResponse.Services.Blind.ServiceCapability = atm.ServicesBlindServiceCapability
	atmResponse.Services.Blind.ServiceActivity = atm.ServicesBlindServiceActivity

	atmResponse.Services.NFCForBankCards.ServiceCapability = atm.ServicesNFCForBankCardsServiceCapability
	atmResponse.Services.NFCForBankCards.ServiceActivity = atm.ServicesNFCForBankCardsServiceActivity

	atmResponse.Services.QRRead.ServiceCapability = atm.ServicesQRReadServiceCapability
	atmResponse.Services.QRRead.ServiceActivity = atm.ServicesQRReadServiceActivity

	atmResponse.Services.SupportsUSD.ServiceCapability = atm.ServicesSupportsUSDServiceCapability
	atmResponse.Services.SupportsUSD.ServiceActivity = atm.ServicesSupportsUSDServiceActivity

	atmResponse.Services.SupportsChargeRUB.ServiceCapability = atm.ServicesSupportsChargeRUBServiceCapability
	atmResponse.Services.SupportsChargeRUB.ServiceActivity = atm.ServicesSupportsChargeRUBServiceActivity

	atmResponse.Services.SupportsEUR.ServiceCapability = atm.ServicesSupportsEURServiceCapability
	atmResponse.Services.SupportsEUR.ServiceActivity = atm.ServicesSupportsEURServiceActivity

	atmResponse.Services.SupportsRUB.ServiceCapability = atm.ServicesSupportsRUBServiceCapability
	atmResponse.Services.SupportsRUB.ServiceActivity = atm.ServicesSupportsRUBServiceActivity

	c.JSON(http.StatusOK, atmResponse)
}

func SearchAtmByName(repository *repository.Repository, c *gin.Context) {
	name := c.Param("query")
	atms, err := repository.GetAtmByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(name) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "для поиска введите какие-либо данные",
		})
	}

	var atmResponses []model.ATMResponse
	for _, atm := range atms {
		atmResponse := model.ATMResponse{
			Address:   atm.Address,
			Latitude:  atm.Latitude,
			Longitude: atm.Longitude,
			AllDay:    atm.AllDay == "true",
		}

		// Маппинг данных из Services
		atmResponse.Services.Wheelchair.ServiceCapability = atm.ServicesWheelchairServiceCapability
		atmResponse.Services.Wheelchair.ServiceActivity = atm.ServicesWheelchairServiceActivity

		atmResponse.Services.Blind.ServiceCapability = atm.ServicesBlindServiceCapability
		atmResponse.Services.Blind.ServiceActivity = atm.ServicesBlindServiceActivity

		atmResponse.Services.NFCForBankCards.ServiceCapability = atm.ServicesNFCForBankCardsServiceCapability
		atmResponse.Services.NFCForBankCards.ServiceActivity = atm.ServicesNFCForBankCardsServiceActivity

		atmResponse.Services.QRRead.ServiceCapability = atm.ServicesQRReadServiceCapability
		atmResponse.Services.QRRead.ServiceActivity = atm.ServicesQRReadServiceActivity

		atmResponse.Services.SupportsUSD.ServiceCapability = atm.ServicesSupportsUSDServiceCapability
		atmResponse.Services.SupportsUSD.ServiceActivity = atm.ServicesSupportsUSDServiceActivity

		atmResponse.Services.SupportsChargeRUB.ServiceCapability = atm.ServicesSupportsChargeRUBServiceCapability
		atmResponse.Services.SupportsChargeRUB.ServiceActivity = atm.ServicesSupportsChargeRUBServiceActivity

		atmResponse.Services.SupportsEUR.ServiceCapability = atm.ServicesSupportsEURServiceCapability
		atmResponse.Services.SupportsEUR.ServiceActivity = atm.ServicesSupportsEURServiceActivity

		atmResponse.Services.SupportsRUB.ServiceCapability = atm.ServicesSupportsRUBServiceCapability
		atmResponse.Services.SupportsRUB.ServiceActivity = atm.ServicesSupportsRUBServiceActivity

		atmResponses = append(atmResponses, atmResponse)
	}

	c.JSON(http.StatusOK, atmResponses)
}
