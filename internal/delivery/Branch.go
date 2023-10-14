package delivery

import (
	"net/http"

	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
	"github.com/SicParv1sMagna/mdhh_backend/internal/pkg/middleware/decode"
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
		openHours, err := decode.UnmarshalOpenHours(branch.OpenHours)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		openHoursIndividual, err := decode.UnmarshalOpenHours(branch.OpenHoursIndividual)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

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

func GetBranchBySearch(repository *repository.Repository, c *gin.Context) {
	var branches []model.Branch

	search := c.Param("query")

	branches, err := repository.GetBranchBySearch(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	var branchResponses []model.BranchResponse

	for _, branch := range branches {
		openHours, err := decode.UnmarshalOpenHours(branch.OpenHours)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		openHoursIndividual, err := decode.UnmarshalOpenHours(branch.OpenHoursIndividual)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

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

func GetBranchById(repository *repository.Repository, c *gin.Context) {

}
