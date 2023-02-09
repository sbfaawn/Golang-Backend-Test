package router

import (
	"Golang-Backend-Test/handler"
	initializers "Golang-Backend-Test/initializer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectToDB(&config)
	initializers.ConnectToRedis(&config)
}

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router := e.Group("/api")
	router.GET("/health-check", handler.HealthCheckHandler)
	// create endpoint
	router.POST("/create-user", handler.CreateUser)
	router.POST("/create-order-item", handler.CreateOrderItem)
	router.POST("/create-order-history", handler.CreateOrderHistory)
	// list endpoint
	router.GET("/list-user", handler.ListUsers)
	router.GET("/list-order-item", handler.ListOrderItems)
	router.GET("/list-order-history", handler.ListOrderHistories)
	//
	router.GET("/detail-user", handler.DetailUser)
	router.GET("/detail-order-item", handler.DetailOrderItem)
	router.GET("/detail-order-history", handler.DetailOrderHistory)
	//
	router.PUT("/update", handler.UpdateHandler)
	//
	router.DELETE("/delete-user", handler.DeleteUser)
	router.DELETE("/delete-order-item", handler.DeleteOrderItem)
	router.DELETE("/delete-order-history", handler.DeleteOrderHistory)

	return e
}
