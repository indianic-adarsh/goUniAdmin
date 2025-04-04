package admin

import (
	"goUniAdmin/internal/config"
	"goUniAdmin/internal/db"
	"goUniAdmin/internal/modules"
	"goUniAdmin/internal/services/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

// adminModule implements the Module interface
type adminModule struct {
	handler *AdminHandler
}

// RegisterRoutes sets up the admin routes
func (m *adminModule) RegisterRoutes(group *gin.RouterGroup, cfg *config.Config, db *db.DB) {
	adminGroup := group.Group("/admins")
	adminGroup.POST("", m.handler.CreateAdmin)
	adminGroup.GET("", m.handler.ListAdmins)
	adminGroup.POST("/login", m.handler.AdminLogin)

	// Protected routes with JWT authentication
	adminGroup.Use(middleware.AuthMiddleware(cfg))
	adminGroup.GET("/:id", m.handler.GetAdmin)
	adminGroup.PUT("/:id", m.handler.UpdateAdmin)
	adminGroup.DELETE("/:id", m.handler.DeleteAdmin)
	adminGroup.GET("/profile", m.handler.GetProfile)
}

// RegisterAdminModule registers the admin module with the given dependencies
func RegisterAdminModule(cfg *config.Config, db *db.DB) {
	service := NewAdminService(db, cfg)
	handler := NewAdminHandler(service)
	modules.RegisterModule(&adminModule{handler: handler})
	log.Println("Admin module registered")
}
