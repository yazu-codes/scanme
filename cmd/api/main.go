package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/yazu-codes/scanme.git/internal/database"
	"github.com/yazu-codes/scanme.git/internal/handlers"
	"github.com/yazu-codes/scanme.git/internal/middleware"
	"github.com/yazu-codes/scanme.git/internal/model"
	"github.com/yazu-codes/scanme.git/internal/service"

	"github.com/spf13/viper"
)

func main() {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)
	logger = logger.With(slog.String("component", "menu_service_backend"))

	config := os.Getenv("CONFIG_YAML")
	if config != "" {
		fmt.Println("CONFIG_YAML environment variable is set. Writing to config.yaml.")
		err := os.WriteFile("./cmd/api/configs/config.yaml", []byte(config), 0600)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("CONFIG_YAML environment variable is not set. Using existing config.yaml.")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./cmd/api/configs") // current directory

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	// -----------------------
	// Extract values
	// -----------------------
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	sslmode := viper.GetString("database.sslmode")
	timezone := viper.GetString("database.timezone")

	jwtSecret := viper.GetString("jwt_secret")

	translationServiceUrl := viper.GetString("translation_service.url")

	// -----------------------
	// Build DSN
	// -----------------------
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:3000"}, // React
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true, // Allow all origins for testing purposes
	}))

	if dsn == "" {
		fmt.Println("DATABASE_URL is not set. Using default SQLite db.")
		dsn = "app.db"
	}

	fmt.Println("DSN:", dsn)

	db := database.Connect(dsn)

	fmt.Println("Auto-migrating database schema...")
	if err := db.AutoMigrate(
		// &models.User{},
		// &models.Post{},
		&model.Menu{},
		&model.MenuConfiguration{},
		&model.MenuItem{},
		&model.MenuOwner{},
		&model.CardMenuCode{},
	); err != nil {
		log.Fatal(err)
	}

	httpClient := http.Client{}

	menuService := service.NewMenuService(db)
	cardMenuCodeService := service.NewCardMenuCodeService(db)
	translationService := service.NewTranslationService(translationServiceUrl, logger)
	publicHandler := handlers.NewPublicHandler(translationService, &httpClient, menuService, cardMenuCodeService)

	router.GET("/yumm", publicHandler.GetYummBrief)
	router.GET("/:name", publicHandler.GetMenuByName)
	router.GET("/c/:code", publicHandler.GetMenuNameByCode)

	// Protected routes
	api := router.Group("/api")

	// Public routes
	api.GET("/", publicHandler.Home)
	// MOVING THIS ONE TO ADMIN api.GET("/menus", publicHandler.GetMenus)
	api.GET("/codes", publicHandler.GetAllCodes)
	api.POST("/login", publicHandler.Login)
	api.POST("/create-menu", publicHandler.CreateMenu)
	api.POST("/create-code", publicHandler.CreateCardMenuCode)
	api.PUT("/update-code", publicHandler.UpdateCardMenuCode)
	api.POST("/create-owner", publicHandler.AddMenuOwner)
	api.PUT("/update-menu", middleware.WithTimeout(2*time.Minute), publicHandler.UpdateMenu)
	api.POST("/suspend-menu/:id", publicHandler.SuspendMenuById)
	api.POST("/yumm-enable/:id", publicHandler.EnableYummById)
	api.POST("/yumm-disable/:id", publicHandler.DisableYummById)
	api.POST("/enable-menu/:id", publicHandler.EnableMenuById)
	api.DELETE("/delete-menu/:id", publicHandler.DeleteMenuById)

	api.Use(middleware.AuthMiddleware(jwtSecret), middleware.RequireRole("admin", "user"))
	{
		api.GET("/menus", publicHandler.GetMenus)
	}

	api.Use(middleware.AuthMiddleware(jwtSecret), middleware.RequireRole("admin"))
	{
		api.GET("/menu-associations/:userId", publicHandler.MenuAssociations)
		api.POST("/menu-associations", publicHandler.SetMenuAssociations)
		api.GET("/profile", handlers.Profile)
		api.GET("/settings", handlers.Settings)
	}

	router.Run(":8080")
}
