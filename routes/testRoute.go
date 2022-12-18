package routes

import (
	"github.com/Gymkhana-Forms/backend/controllers"
	"github.com/gin-gonic/gin"
)

func TestRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/test", controllers.Test())
}
