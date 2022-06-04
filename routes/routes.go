package routes

import (
	"github.com/SherzodAbdullajonov/e-wallet/controller"
	"github.com/gin-gonic/gin"
)

func UseRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.POST("/balance", controller.GetBalance())
	incomingRoutes.POST("/top_up", controller.GetTotalTopUps())
	incomingRoutes.POST("/user", controller.CheckAcount())
	incomingRoutes.POST("/total", controller.GetTotalTopUps())
}