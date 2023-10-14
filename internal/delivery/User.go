package delivery

import (
	"fmt"
	"net/http"

	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
	email "github.com/SicParv1sMagna/mdhh_backend/internal/pkg/middleware/emailConfirmation"
	"github.com/SicParv1sMagna/mdhh_backend/internal/pkg/middleware/password"
	"github.com/SicParv1sMagna/mdhh_backend/internal/pkg/middleware/validators"
	"github.com/SicParv1sMagna/mdhh_backend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func RegisterUser(repository *repository.Repository, c *gin.Context, s *email.EmailSender) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(user)
	if err := validators.ValidateRegistrationData(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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

	uniqueCode := email.GenerateUniqueCode()

	user.AccessToken = uniqueCode

	fmt.Println(uniqueCode, user.Email)
	err = s.SendConfirmationEmail(uniqueCode, user.Email)
	if err != nil {
		fmt.Println(err)
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

func ConfirmRegistration(repository *repository.Repository, c *gin.Context) {
	var jsonData map[string]interface{}
	if err := c.BindJSON(&jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	confirmationCode, ok := jsonData["confirmationCode"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "неверный или отсутствующий код подтверждения",
		})
		return
	}

	candidate, err := repository.GetUserByToken(confirmationCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ошибка при попытке определить код",
		})
		return
	}

	err = repository.ConfirmRegistration(candidate.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "регистрация подтверждена",
	})
}

func AuthUser(repository *repository.Repository, store *sessions.CookieStore, c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(user)
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

	c.JSON(http.StatusOK, gin.H{
		"message": "авторизован",
	})
}

func ResendConfirmationCode(repository *repository.Repository, c *gin.Context, s *email.EmailSender) {
	var jsonData map[string]interface{}
	if err := c.BindJSON(&jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	uEmail, ok := jsonData["Email"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "неверный или отсутствующий код подтверждения",
		})
		return
	}

	uniqueCode := email.GenerateUniqueCode()

	err := s.SendConfirmationEmail(uniqueCode, uEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	err = repository.UpdateUserAccessToken(uEmail, uniqueCode)
	if err != nil {
		c.JSON(http.StatusOK, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "проверьте вашу почту",
	})
}
