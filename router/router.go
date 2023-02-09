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
	router.GET("/list", handler.ListHandler)
	router.GET("/detail", handler.DetailHandler)
	router.POST("/create-user", handler.CreateUser)
	router.POST("/create-order-item", handler.CreateOrderItem)
	router.POST("/create-order-history", handler.CreateOrderHistory)
	router.PUT("/update", handler.UpdateHandler)
	router.DELETE("/delete", handler.DeleteHandler)

	return e
}
