package handler

import (
	"Golang-Backend-Test/helper"
	"Golang-Backend-Test/model/request"
	response2 "Golang-Backend-Test/model/response"
	"Golang-Backend-Test/service"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func HealthCheckHandler(c echo.Context) error {
	response := response2.Response{Message: "Project is Up"}
	return c.JSON(http.StatusOK, response)
}

// /
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

///

func ListUsers(c echo.Context) error {
	users, err := service.GetUsers()
	if err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to Retrieve all data",
		Data:    users,
	}

	return c.JSON(http.StatusOK, response)
}

func ListOrderItems(c echo.Context) error {
	orderItems, err := service.GetOrderItems()
	if err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to Retrieve all data",
		Data:    orderItems,
	}

	return c.JSON(http.StatusOK, response)
}

func ListOrderHistories(c echo.Context) error {
	orderHistories, err := service.GetOrderHistories()
	if err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to Retrieve all data",
		Data:    orderHistories,
	}

	return c.JSON(http.StatusOK, response)
}

// /
func DetailUser(c echo.Context) error {
	paramId := c.QueryParam("id")
	idInt, err := strconv.Atoi(paramId)
	if err != nil {
		return err
	}

	user, err := service.GetUserById(&idInt)
	if err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to Retrieve data with Id " + paramId,
		Data:    user,
	}

	return c.JSON(http.StatusOK, response)
}

func DetailOrderItem(c echo.Context) error {
	paramId := c.QueryParam("id")
	idInt, err := strconv.Atoi(paramId)
	if err != nil {
		return err
	}

	orderItem, err := service.GetOrderItemById(&idInt)
	if err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to Retrieve data with Id " + paramId,
		Data:    orderItem,
	}

	return c.JSON(http.StatusOK, response)
}

func DetailOrderHistory(c echo.Context) error {
	paramId := c.QueryParam("id")
	idInt, err := strconv.Atoi(paramId)
	if err != nil {
		return err
	}

	orderHistory, err := service.GetOrderHistoriesById(&idInt)
	if err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to Retrieve data with Id " + paramId,
		Data:    orderHistory,
	}

	return c.JSON(http.StatusOK, response)
}

// /
func UpdateUser(c echo.Context) error {
	o := new(request.User)
	if err := c.Bind(o); err != nil {
		return err
	}

	if o.Id == nil {
		return errors.New("Id need to be specified in request")
	}

	user, err := service.UpdateUser(o)
	if err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to update data with Id " + strconv.Itoa(*o.Id),
		Data:    user,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateOrderItem(c echo.Context) error {
	o := new(request.OrderItem)
	if err := c.Bind(o); err != nil {
		return err
	}

	if o.Id == nil {
		return errors.New("Id need to be specified in request")
	}

	orderItem, err := service.UpdateOrderItem(o)
	if err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to update data with Id " + strconv.Itoa(*o.Id),
		Data:    orderItem,
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateOrderHistory(c echo.Context) error {
	o := new(request.OrderHistory)
	if err := c.Bind(o); err != nil {
		return err
	}

	if o.Id == nil {
		return errors.New("Id need to be specified in request")
	}

	orderHistory, err := service.UpdateOrderHistory(o)
	if err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to update data with Id " + strconv.Itoa(*o.Id),
		Data:    orderHistory,
	}
	return c.JSON(http.StatusOK, response)
}

// /
func DeleteUser(c echo.Context) error {
	paramId := c.QueryParam("id")
	idInt, err := strconv.Atoi(paramId)
	if err != nil {
		return err
	}

	err = service.SoftDeleteUser(&idInt)
	if err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to Delete data with Id " + paramId,
		Data:    nil,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteOrderItem(c echo.Context) error {
	paramId := c.QueryParam("id")
	idInt, err := strconv.Atoi(paramId)
	if err != nil {
		return err
	}

	err = service.SoftDeleteOrderItem(&idInt)
	if err != nil {
		return err
	}

	response := response2.Response{
		Message: "Success to Delete data with Id " + paramId,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, response)
}
