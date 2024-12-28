package main

import (
	"Marketplace-backend/internal/framework/db"
	"Marketplace-backend/internal/interface_adopter/controller"
	"Marketplace-backend/internal/interface_adopter/gateway"
	"Marketplace-backend/internal/interface_adopter/routes"
	"Marketplace-backend/internal/service"
	"Marketplace-backend/pkg/config"
	"log"

	_ "github.com/lib/pq" // This imports the PostgreSQL driver

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env (it will loof for .env in the root directory)
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning .env file not found. using enviroment variables")
	}
}

func main() {
	// Load the database configuration  from enviroment variables are .env
	dbConfig := config.LoadDBConfig()

	// debug: print the loaded database configuration
	log.Printf("DB Config: Host=%s, Port=%s, User=%s, Password=%s, DBName=%s, SSLMode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)

	// connection to the database
	database, err := db.ConnectDB(dbConfig)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer database.Close()

	// Create the required tables if they don't exist
	if err := db.CreatTables(database); err != nil {
		log.Fatal("Error creating tables:", err)
	}

	// Intialize the repository
	userRepository := gateway.NewUserRepository(database)
	tokenRepository := gateway.NewTokenRepository(database)
	productRepository := gateway.NewProductRepositoryImpl(database)
	salesRepository := gateway.NewSalesRepository(database)
	reviewRatingRepository := gateway.NewReviewRatingRepositoryImpl(database)
	orderRepository := gateway.NewOrderRepository(database)

	//Intialize the service
	userService := service.NewUserService(userRepository, tokenRepository)
	productService := service.NewProductService(productRepository, tokenRepository)
	saleservice := service.NewSalesService(salesRepository, tokenRepository)
	reviewRatingService := service.NewReviewRatingService(reviewRatingRepository, tokenRepository)
	orderService := service.NewOrderService(orderRepository, tokenRepository)

	// Initalize controller
	userController := controller.NewUserController(userService)
	productController := controller.NewProductController(productService)
	salesController := controller.NewSalesContriller(saleservice)
	reviewRatingController := controller.NewReviewRatingController(reviewRatingService)
	orderController := controller.NewOrderController(orderService)

	// Intialize Gin router
	r := gin.Default()

	// Apply CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://your-frontend-domain.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Rigester user-raleted routes with the token repository for middleware
	routes.RegisterUserRoutes(r, userController, tokenRepository)
	routes.RegisterProductRoutes(r, *productController, tokenRepository)
	routes.RegisterSalesRoutes(r, *salesController, tokenRepository)
	routes.RegisterReviewRatingRoutes(r, *reviewRatingController, tokenRepository)
	routes.RegisterOrderRoutes(r, *orderController, tokenRepository)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error running the server: ", err)
	}
}
