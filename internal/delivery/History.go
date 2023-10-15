package delivery

import (
	"net/http"

	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
	"github.com/SicParv1sMagna/mdhh_backend/internal/repository"
	"github.com/gin-gonic/gin"
)

func AddRoute(repository *repository.Repository, c *gin.Context) {
	var route model.Route

	if err := c.BindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := repository.AddRoute(&route); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"id":     route.ID,
	})
}

func GetAllRoutes(repository *repository.Repository, c *gin.Context) {
	routes, err := repository.GetAllRoutes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, routes)
}
