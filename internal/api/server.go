package api

import (
	"log"

	"github.com/SicParv1sMagna/mdhh_backend/internal/delivery"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func (a *Application) StartServer() {
	router := gin.Default()

	store := sessions.NewCookieStore([]byte("SuperSecretKey"))

	user := router.Group("/user")
	{
		//	localhost:8080/user/register
		user.POST("/register", func(ctx *gin.Context) {
			delivery.RegisterUser(a.repository, ctx)
		})

		//	localhost:8080/user/login
		user.POST("/login", func(ctx *gin.Context) {
			delivery.AuthUser(a.repository, store, ctx)
		})
	}

	err := router.Run()
	if err != nil {
		log.Fatalf("error, while running the server")
	}
}
