package delivery

import (
	"errors"
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
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(user)
	if err := validators.ValidateRegistrationData(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		fmt.Println(err)
		return
	}

	candidate, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	uniqueCode := email.GenerateUniqueCode()

	user.AccessToken = uniqueCode

	fmt.Println(user.Email)
	err = s.SendConfirmationEmail(uniqueCode, user.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "пользователь зарегестрирован",
	})
}

func ConfirmRegistration(repository *repository.Repository, c *gin.Context) {
	var jsonData map[string]interface{}
	if err := c.BindJSON(&jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "регистрация подтверждена",
	})
}

func AuthUser(repository *repository.Repository, store *sessions.CookieStore, c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := validators.ValidateAuthorizationData(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	candidate, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !candidate.IsConfirmed {
		c.JSON(http.StatusInternalServerError, errors.New("аккаунт не подтвержден"))
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
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 3,
		HttpOnly: true,
		//SameSite: http.SameSiteNoneMode,
	}

	session.Values["userID"] = candidate.User_ID

	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "авторизован",
	})
}

func LogoutUser(repository *repository.Repository, store *sessions.CookieStore, c *gin.Context) {
	session, err := store.Get(c.Request, "J_SESSION")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	session.Values["userID"] = nil

	session.Options.MaxAge = -1

	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "вышел",
	})
}

func ResendConfirmationCode(repository *repository.Repository, c *gin.Context, s *email.EmailSender) {
	var jsonData map[string]interface{}
	if err := c.BindJSON(&jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = repository.UpdateUserAccessToken(uEmail, uniqueCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "проверьте вашу почту",
	})
}

func GetUserById(repository *repository.Repository, store *sessions.CookieStore, c *gin.Context) {
	session, err := store.Get(c.Request, "J_SESSION")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userID := session.Values["userID"]
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "пользователь не авторизован",
		})
		return
	}

	id, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("неверный идентифиактор пользователь").Error())
		return
	}

	if id < 1 {
		c.JSON(http.StatusBadRequest, errors.New("id пользователя не может быть отрицательным").Error())
		return
	}

	user, err := repository.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func EditUserData(repository *repository.Repository, store *sessions.CookieStore, c *gin.Context) {
	session, err := store.Get(c.Request, "J_SESSION")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userID := session.Values["userID"]
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Пользователь не авторизован",
		})
		return
	}

	id, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("неверный идентификатор пользователя").Error())
		return
	}

	if id < 1 {
		c.JSON(http.StatusBadRequest, errors.New("id пользователя не может быть отрицательным").Error)
		return
	}

	var jsonData map[string]interface{}
	if err := c.BindJSON(&jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	newFirstName, ok := jsonData["firstname"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, errors.New("ошибка при получении имени").Error())
		return
	}

	newSecondName, ok := jsonData["secondname"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, errors.New("ошибка при получении фамилии").Error())
		return
	}

	newEmail, ok := jsonData["email"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, errors.New("ошибка при получении email").Error())
		return
	}

	newLegalEntity, ok := jsonData["legalentity"].(bool)
	if !ok {
		c.JSON(http.StatusInternalServerError, errors.New("ошибка при получении вашего статуса").Error())
		return
	}

	user := &model.User{
		FirstName:   newFirstName,
		SecondName:  newSecondName,
		Email:       newEmail,
		LegalEntity: newLegalEntity,
	}

	err = repository.UpdateUserData(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "данные обновлены",
	})
}

func DeleteUser(repository *repository.Repository, store *sessions.CookieStore, c *gin.Context) {
	session, err := store.Get(c.Request, "J_SESSION")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userID, ok := session.Values["userID"]
	if !ok {
		c.JSON(http.StatusUnauthorized, "пользователь не авторизован")
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, "неверный формат идентификатора пользователя")
		return
	}

	session.Options.MaxAge = -1
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = repository.DeleteUser(userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "удален",
	})
}
