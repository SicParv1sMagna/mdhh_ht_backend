package delivery

import (
	"net/http"

	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
	"github.com/SicParv1sMagna/mdhh_backend/internal/pkg/middleware/password"
	"github.com/SicParv1sMagna/mdhh_backend/internal/pkg/middleware/validators"
	"github.com/SicParv1sMagna/mdhh_backend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func RegisterUser(repository *repository.Repository, c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := validators.ValidateRegistrationData(user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	candidate, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if candidate == user {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "такой пользователь уже существует",
		})
		return
	}

	user.Password, err = password.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "пользователь зарегестрирован",
	})
}

func AuthUser(repository *repository.Repository, store *sessions.CookieStore, c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := validators.ValidateAuthorizationData(user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	candidate, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if ok := password.CheckPasswordHash(user.Password, candidate.Password); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "пароли не совпадают",
		})
		return
	}

	session, err := store.Get(c.Request, "J_SESSION")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 3,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	session.Values["userID"] = candidate.User_ID

	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, session)

	c.JSON(http.StatusOK, gin.H{
		"message": "авторизован",
	})
}
