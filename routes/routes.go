package routes

import (
	"yanwr/digital-bank/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IRoutes interface {
	LoadRoutes(router *gin.Engine) *gin.Engine
}

type Routes struct {
	accountController controllers.IAccountController
}

func NewRoutes(conDB *gorm.DB) IRoutes {
	return &Routes{
		accountController: controllers.NewAccountController(conDB),
	}
}

func (r *Routes) LoadRoutes(router *gin.Engine) *gin.Engine {
	router = ConfigCORS(router)
	main := router.Group("app")
	{
		accounts := main.Group("accounts")
		{
			accounts.POST("/", r.accountController.CreateAccount)
			accounts.GET("/", r.accountController.IndexAllAccounts)
			accounts.GET("/:account_id/balance", r.accountController.ShowBalanceAccount)
		}
		// transfers := main.Group("transfers")
		// {
		// 	transfers.GET("/", controllers.ShowTransfersFromCurrentUser)
		// 	transfers.POST("/", controllers.CreateTransfersTo)
		// }
	}
	return router
}

func ConfigCORS(router *gin.Engine) *gin.Engine {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080"
		},
	}))
	return router
}
