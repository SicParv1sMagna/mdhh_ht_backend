package api

import (
	"github.com/SicParv1sMagna/mdhh_backend/internal/pkg/middleware/cors"
	"log"

	"github.com/SicParv1sMagna/mdhh_backend/internal/delivery"
	emailsender "github.com/SicParv1sMagna/mdhh_backend/internal/pkg/middleware/emailConfirmation"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func (a *Application) StartServer() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("SuperSecretKey"))

	// Настройка CORS
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost:3000"} // Добавьте адрес клиента
	//config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	//config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	//config.AllowCredentials = true
	//router.Use(cors.New(config))

	router.Use(cors.CORSMiddleware())

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
	//api.Use(auth.AuthCheck(store))
	{
		user := api.Group("/users")
		{
			user.POST("/logout", func(ctx *gin.Context) {
				delivery.LogoutUser(a.repository, store, ctx)
			})

			user.POST("/get-user-by-id", func(ctx *gin.Context) {
				delivery.GetUserById(a.repository, store, ctx)
			})

			user.PUT("/edit-user-data", func(ctx *gin.Context) {
				delivery.EditUserData(a.repository, store, ctx)
			})

			user.DELETE("/delete-user", func(ctx *gin.Context) {
				delivery.DeleteUser(a.repository, store, ctx)
			})
		}

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

			//	http://localhost:8080/api/branches/get-branch-by-id
			branches.GET("/get-branch-by-id/:id", func(ctx *gin.Context) {
				delivery.GetBranchById(a.repository, ctx)
			})

			//	http://localhost:8080/api/branches/get-nearest-branches-with-talons?latitude&&longitude
			branches.GET("/ws/get-nearest-branches-with-talons", func(ctx *gin.Context) {
				delivery.GetBranchesWithTalons(a.repository, ctx)
			})
		}

		atm := api.Group("/atm")
		{
			// /api/atm/get-all-atm
			atm.GET("/get-all-atm", func(ctx *gin.Context) {
				delivery.GetAllAtms(a.repository, ctx)
			})

			atm.GET("/get-atm-by-id/:id", func(ctx *gin.Context) {
				delivery.GetAtmById(a.repository, ctx)
			})

			atm.GET("/search-atm-by-name/:query", func(ctx *gin.Context) {
				delivery.SearchAtmByName(a.repository, ctx)
			})
		}
	}

	moderator := api.Group("/moderator")
	{
		moderator.POST("/talon", func(ctx *gin.Context) {
			delivery.AddTalon(a.repository, ctx)
		})

		moderator.DELETE("/talon", func(ctx *gin.Context) {
			delivery.DeleteTalon(a.repository, ctx)
		})
	}

	err = router.Run()
	if err != nil {
		log.Fatalf("error, while running the server")
	}
}
