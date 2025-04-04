package main

import (
	"log"
	"time"

	_ "goUniAdmin/docs"
	"goUniAdmin/internal/config"
	"goUniAdmin/internal/db"
	"goUniAdmin/internal/modules"
	"goUniAdmin/internal/modules/admin"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Uni Admin API
// @version 1.0
// @description Admin panel API for Go Uni Admin
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@gouniadmin.com
// @license.name MIT
// @host localhost:5000
// @BasePath /api
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.NewDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	config.ConfigInstance = cfg
	db.DBInstance = dbConn

	admin.RegisterAdminModule(cfg, dbConn)

	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API routes group

	modules.InitializeModules(router, cfg, dbConn)

	// Swagger setup
	router.GET("/swagger/*any", func(c *gin.Context) {
		if cfg.IsHTTPAuthForSwagger {
			gin.BasicAuth(gin.Accounts{
				cfg.SwaggerAuthUser: cfg.SwaggerAuthPassword,
			})(c)
			if c.IsAborted() {
				return
			}
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	})

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 404 handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"success": false,
			"error":   "Endpoint not found",
			"path":    c.Request.URL.Path,
		})
	})

	log.Printf("Starting server on %s", cfg.Port)
	log.Fatal(router.Run(cfg.Port))
}
