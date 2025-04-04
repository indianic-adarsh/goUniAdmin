package modules

import (
	"goUniAdmin/internal/config"
	"goUniAdmin/internal/db"

	"github.com/gin-gonic/gin"
)

// Module defines the interface for all modules
type Module interface {
	RegisterRoutes(group *gin.RouterGroup, cfg *config.Config, db *db.DB)
}

// modules holds all registered modules
var modules []Module

// RegisterModule adds a module to the registry
func RegisterModule(m Module) {
	modules = append(modules, m)
}

// Modules returns the list of registered modules (for debugging)
func Modules() []Module {
	return modules
}

// InitializeModules initializes all registered modules under /api prefix
func InitializeModules(router *gin.Engine, cfg *config.Config, db *db.DB) {
	apiGroup := router.Group("/api")
	{
		for _, m := range modules {
			m.RegisterRoutes(apiGroup, cfg, db)
		}
	}
}
