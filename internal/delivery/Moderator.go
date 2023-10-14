package delivery

import (
	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
	"github.com/SicParv1sMagna/mdhh_backend/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//func getFailJson(msg string) ([]byte, error) {
//	str := `{"status":  "fail", "message": ` + msg + `}`
//
//	if jsonString, err := json.Marshal(str); err != nil {
//		return []byte{}, err
//	} else {
//		return jsonString, nil
//	}
//}

func AddTalon(repository *repository.Repository, c *gin.Context) {
	var talon = model.Talon{}

	var err error
	if err = c.ShouldBindJSON(&talon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err = repository.AddTalon(&talon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"id":     talon.ID,
	})

}

func DeleteTalon(repository *repository.Repository, c *gin.Context) {
	id, err := strconv.Atoi(c.DefaultQuery("id", ""))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if id < 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err = repository.DeleteTalon(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
