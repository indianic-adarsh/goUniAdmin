package admin

import (
	"goUniAdmin/internal/config"
	"goUniAdmin/internal/services/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the admin routes
func RegisterRoutes(router *gin.Engine, handler *AdminHandler, cfg *config.Config) {
	adminGroup := router.Group("/admins")
	adminGroup.POST("", handler.CreateAdmin)
	adminGroup.GET("", handler.ListAdmins)
	adminGroup.POST("/login", handler.AdminLogin)

	// Protected routes with JWT authentication
	adminGroup.Use(middleware.AuthMiddleware(cfg))
	adminGroup.GET("/:id", handler.GetAdmin)
	adminGroup.PUT("/:id", handler.UpdateAdmin)
	adminGroup.DELETE("/:id", handler.DeleteAdmin)
	adminGroup.GET("/profile", handler.GetProfile)
}
