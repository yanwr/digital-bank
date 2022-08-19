package routes

import (
	"yanwr/digital-bank/controllers"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("digital-bank")
	{
		accounts := main.Group("accounts")
		{
			accounts.GET("/", controllers.ShowAllAccounts)
			accounts.GET("/:account_id/balance", controllers.IndexBalanceAccount)
		}

		transfers := main.Group("transfers")
		{
			transfers.GET("/", controllers.ShowTransfersFromCurrentUser)
			transfers.POST("/", controllers.CreateTransfersTo)
		}
	}
	return router
}
