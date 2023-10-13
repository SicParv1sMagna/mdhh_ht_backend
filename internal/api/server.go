package api

import (
	"log"

	"github.com/SicParv1sMagna/mdhh_backend/internal/delivery"
	emailsender "github.com/SicParv1sMagna/mdhh_backend/internal/pkg/middleware/emailConfirmation"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func (a *Application) StartServer() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("SuperSecretKey"))

	// Настройка CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost"}
	router.Use(cors.New(config))

	sender, err := emailsender.New()
	if err != nil {
		log.Fatal(err)
	}

	user := router.Group("/user")
	{
		// http://localhost:8080/user/register
		user.POST("/register", func(ctx *gin.Context) {
			delivery.RegisterUser(a.repository, ctx, sender)
		})

		// http://localhost:8080/user/login
		user.POST("/login", func(ctx *gin.Context) {
			delivery.AuthUser(a.repository, store, ctx)
		})

		//	http://localhost:8080/user/confirm-registration
		user.POST("/confirm-registration", func(ctx *gin.Context) {
			delivery.ConfirmRegistration(a.repository, ctx)
		})

		//	http://localhost:8080/user/resend-code
		user.PUT("/resend-code", func(ctx *gin.Context) {
			delivery.ResendConfirmationCode(a.repository, ctx, sender)
		})
	}

	err = router.Run(":80")
	if err != nil {
		log.Fatalf("error, while running the server")
	}
}
