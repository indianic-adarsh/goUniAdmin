package main

import (
	"goUniAdmin/internal/config"
	"goUniAdmin/internal/db"
	"goUniAdmin/internal/modules/admin"
	"log"

	"net/http"

	_ "goUniAdmin/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title goUniAdmin API
// @version 1.0
// @description Admin panel API for goUniAdmin
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@gouniadmin.com
// @license.name MIT
// @host localhost:5000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize GORM database connection
	dbConn, err := db.NewDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the admin table
	if err := dbConn.AutoMigrate(&admin.Admin{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	router := gin.Default()

	// Set up logging middleware
	router.Use(gin.Logger())

	// Set up CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", cfg.AllowedOrigins)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// Set up Gin mode
	gin.SetMode(cfg.GinMode)

	// Set up route for uploading files

	// Define a GET route
	router.GET("/", func(c *gin.Context) {
		// Respond with JSON data
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// Initialize admin module with database
	adminService := admin.NewAdminService(dbConn, cfg)
	adminHandler := admin.NewAdminHandler(adminService)
	admin.RegisterRoutes(router, adminHandler, cfg)

	// Serve Swagger UI at /swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Starting %s server on %s (Environment: %s)", cfg.AppName, cfg.Port, cfg.Environment)
	log.Fatal(router.Run(cfg.Port))
}
