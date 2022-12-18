package routes

import (
	"github.com/Gymkhana-Forms/backend/controllers"
	"github.com/Gymkhana-Forms/backend/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("admin/forms", controllers.SubmittedForms())
}
