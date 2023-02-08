package main

import (
	initializers "Golang-Backend-Test/initializer"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectToDB(&config)
	initializers.ConnectToRedis(&config)
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router := e.Group("/api")
	router.GET("/health-check", func(c echo.Context) error {
		result := struct {
			Message string `json:"message"`
		}{
			"Project is Up",
		}
		return c.JSON(http.StatusOK, result)
	})

	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
