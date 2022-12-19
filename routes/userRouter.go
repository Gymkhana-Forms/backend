package routes

import (
	"github.com/Gymkhana-Forms/backend/controllers"
	"github.com/Gymkhana-Forms/backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users/forms", controllers.AllFormTemplates())
	incomingRoutes.GET("/users/forms/:form_type", controllers.SelectForm())
	// incomingRoutes.POST("/users/forms/:form_type", controllers.SubmitForm())
	incomingRoutes.GET("/users/status", controllers.GetAllForms())
	incomingRoutes.GET("/users/status/:form_id", controllers.GetForm())
}
