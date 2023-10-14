package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
	"github.com/SicParv1sMagna/mdhh_backend/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetAllBranches(repository *repository.Repository, c *gin.Context) {
	branches, err := repository.GetAllBranches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var branchResponses []model.BranchResponse

	for _, branch := range branches {
		var openHours []model.OpenHoursType
		if err := json.Unmarshal(branch.OpenHours, &openHours); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		var openHoursIndividual []model.OpenHoursType
		if err := json.Unmarshal(branch.OpenHoursIndividual, &openHoursIndividual); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		// Create a new BranchResponse with unmarshaled data
		response := model.BranchResponse{
			Branch_ID:           branch.Branch_ID,
			SalePointName:       branch.SalePointName,
			Address:             branch.Address,
			Status:              branch.Status,
			OpenHours:           openHours,
			RKO:                 branch.RKO,
			OpenHoursIndividual: openHoursIndividual,
			OfficeType:          branch.OfficeType,
			SalePointFormat:     branch.SalePointFormat,
			SUOAvailability:     branch.SUOAvailability,
			HasRamp:             branch.HasRamp,
			Latitude:            branch.Latitude,
			Longitude:           branch.Longitude,
			MetroStation:        branch.MetroStation,
			Distance:            branch.Distance,
			KEP:                 branch.KEP,
			MyBranch:            branch.MyBranch,
			Network:             branch.Network,
			SalePointCode:       branch.SalePointCode,
		}

		branchResponses = append(branchResponses, response)
	}

	c.JSON(http.StatusOK, branchResponses)
}
