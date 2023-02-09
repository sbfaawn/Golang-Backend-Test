package handler

import (
	"Golang-Backend-Test/helper"
	"Golang-Backend-Test/model/request"
	response2 "Golang-Backend-Test/model/response"
	"Golang-Backend-Test/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheckHandler(c echo.Context) error {
	response := response2.Response{Message: "Project is Up"}
	return c.JSON(http.StatusOK, response)
}

func DetailHandler(c echo.Context) error {
	response := response2.Response{Message: "Get Endpoint is Called"}
	return c.JSON(http.StatusOK, response)
}

func ListHandler(c echo.Context) error {
	response := []struct {
		Data   string `json:"data"`
		IsGood bool   `json:"isGood"`
	}{
		{"data1", true},
		{"data2", true},
		{"data3", false},
		{"data4", true},
	}

	return c.JSON(http.StatusOK, response)
}

func CreateUser(c echo.Context) error {
	o := new(request.User)
	if err := c.Bind(o); err != nil {
		return err
	}

	if err := service.InsertUser(o); err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to add data",
		Data:    o,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateOrderItem(c echo.Context) error {
	o := new(request.OrderItem)
	if err := c.Bind(o); err != nil {
		return err
	}

	if err := helper.IsValidDate(o.ExpiredAt); err != nil {
		return err
	}

	if err := service.InsertOrderItem(o); err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to add data",
		Data:    o,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateOrderHistory(c echo.Context) error {
	o := new(request.OrderHistory)
	if err := c.Bind(o); err != nil {
		return err
	}

	err := service.InsertOrderHistory(o)
	if err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to add data",
		Data:    o,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateHandler(c echo.Context) error {
	response := response2.Response{Message: "Update Endpoint is Called"}
	return c.JSON(http.StatusOK, response)
}

func DeleteHandler(c echo.Context) error {
	response := response2.Response{Message: "Delete Endpoint is Called"}
	return c.JSON(http.StatusOK, response)
}
