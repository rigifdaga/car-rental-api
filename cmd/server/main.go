package main

import (
	"car-rental-api/internal/config"
	"car-rental-api/internal/database"
	"car-rental-api/internal/handlers"
	"car-rental-api/internal/repositories"
	"car-rental-api/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repositories
	customerRepo := repositories.NewCustomerRepository(db)
	carRepo := repositories.NewCarRepository(db)
	rentalRepo := repositories.NewRentalRepository(db)
	membershipRepo := repositories.NewMembershipRepository(db)
	bookingTypeRepo := repositories.NewBookingTypeRepository(db)
	driverRepo := repositories.NewDriverRepository(db)
	driverIncentiveRepo := repositories.NewDriverIncentiveRepository(db)

	// Initialize services
	customerService := services.NewCustomerService(customerRepo)
	carService := services.NewCarService(carRepo)
	membershipService := services.NewMembershipService(membershipRepo)
	bookingTypeService := services.NewBookingTypeService(bookingTypeRepo)
	driverService := services.NewDriverService(driverRepo)
	driverIncentiveService := services.NewDriverIncentiveService(driverIncentiveRepo, rentalRepo, carRepo)
	rentalService := services.NewRentalService(rentalRepo, customerRepo, carRepo, bookingTypeRepo, driverRepo, driverIncentiveService, membershipRepo)

	// Initialize handlers
	customerHandler := handlers.NewCustomerHandler(customerService)
	carHandler := handlers.NewCarHandler(carService)
	rentalHandler := handlers.NewRentalHandler(rentalService)
	membershipHandler := handlers.NewMembershipHandler(membershipService)
	bookingTypeHandler := handlers.NewBookingTypeHandler(bookingTypeService)
	driverHandler := handlers.NewDriverHandler(driverService)
	driverIncentiveHandler := handlers.NewDriverIncentiveHandler(driverIncentiveService)

	// Setup routes
	router := setupRoutes(customerHandler, carHandler, rentalHandler, membershipHandler, bookingTypeHandler, driverHandler, driverIncentiveHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupRoutes(
	customerHandler *handlers.CustomerHandler,
	carHandler *handlers.CarHandler,
	rentalHandler *handlers.RentalHandler,
	membershipHandler *handlers.MembershipHandler,
	bookingTypeHandler *handlers.BookingTypeHandler,
	driverHandler *handlers.DriverHandler,
	driverIncentiveHandler *handlers.DriverIncentiveHandler,
) *gin.Engine {
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "Car Rental API is running",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Customer routes
		customers := v1.Group("/customers")
		{
			customers.POST("", customerHandler.CreateCustomer)
			customers.GET("", customerHandler.GetAllCustomers)
			customers.GET("/:id", customerHandler.GetCustomerByID)
			customers.PUT("/:id", customerHandler.UpdateCustomer)
			customers.DELETE("/:id", customerHandler.DeleteCustomer)
		}

		// Car routes
		cars := v1.Group("/cars")
		{
			cars.POST("", carHandler.CreateCar)
			cars.GET("", carHandler.GetAllCars)
			cars.GET("/:id", carHandler.GetCarByID)
			cars.PUT("/:id", carHandler.UpdateCar)
			cars.DELETE("/:id", carHandler.DeleteCar)
			cars.GET("/available", carHandler.GetAvailableCars)
		}

		// Rental routes
		rentals := v1.Group("/rentals")
		{
			rentals.POST("", rentalHandler.CreateRental)
			rentals.GET("", rentalHandler.GetAllRentals)
			rentals.GET("/:id", rentalHandler.GetRentalByID)
			rentals.PUT("/:id", rentalHandler.UpdateRental)
			rentals.DELETE("/:id", rentalHandler.DeleteRental)
			rentals.GET("/active", rentalHandler.GetActiveRentals)
		}

		// Membership routes
		memberships := v1.Group("/memberships")
		{
			memberships.POST("", membershipHandler.CreateMembership)
			memberships.GET("", membershipHandler.GetAllMemberships)
			memberships.GET("/:id", membershipHandler.GetMembershipByID)
			memberships.PUT("/:id", membershipHandler.UpdateMembership)
			memberships.DELETE("/:id", membershipHandler.DeleteMembership)
		}

		// Booking Type routes
		bookingTypes := v1.Group("/booking-types")
		{
			bookingTypes.POST("", bookingTypeHandler.CreateBookingType)
			bookingTypes.GET("", bookingTypeHandler.GetAllBookingTypes)
			bookingTypes.GET("/:id", bookingTypeHandler.GetBookingTypeByID)
			bookingTypes.PUT("/:id", bookingTypeHandler.UpdateBookingType)
			bookingTypes.DELETE("/:id", bookingTypeHandler.DeleteBookingType)
		}

		// Driver routes
		drivers := v1.Group("/drivers")
		{
			drivers.POST("", driverHandler.CreateDriver)
			drivers.GET("", driverHandler.GetAllDrivers)
			drivers.GET("/:id", driverHandler.GetDriverByID)
			drivers.PUT("/:id", driverHandler.UpdateDriver)
			drivers.DELETE("/:id", driverHandler.DeleteDriver)
		}

		// Driver Incentive routes
		driverIncentives := v1.Group("/driver-incentives")
		{
			driverIncentives.POST("", driverIncentiveHandler.CreateDriverIncentive)
			driverIncentives.GET("", driverIncentiveHandler.GetAllDriverIncentives)
			driverIncentives.GET("/:id", driverIncentiveHandler.GetDriverIncentiveByID)
			driverIncentives.PUT("/:id", driverIncentiveHandler.UpdateDriverIncentive)
			driverIncentives.DELETE("/:id", driverIncentiveHandler.DeleteDriverIncentive)
		}
	}

	return router
}
