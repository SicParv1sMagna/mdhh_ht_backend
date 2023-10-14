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

	api := router.Group("/api")
	{
		branches := api.Group("/branches")
		{
			//	http://localhost:8080/api/branches/get-all-branches
			branches.GET("/get-all-branches", func(ctx *gin.Context) {
				delivery.GetAllBranches(a.repository, ctx)
			})

			//	http://localhost:8080/api/branches/get-branch-by-search
			branches.GET("/get-branch-by-search/:query", func(ctx *gin.Context) {
				delivery.GetBranchBySearch(a.repository, ctx)
			})

			branches.GET("/get-branch-by-id/:id", func(ctx *gin.Context) {
				delivery.GetBranchById(a.repository, ctx)
			})
		}

		moderator := router.Group("/moderator")
		{
			moderator.POST("/talon", func(ctx *gin.Context) {
				delivery.AddTalon(a.repository, ctx)
			})

			moderator.DELETE("/talon", func(ctx *gin.Context) {
				delivery.DeleteTalon(a.repository, ctx)
			})
		}
	}

	err = router.Run()
	if err != nil {
		log.Fatalf("error, while running the server")
	}
}
