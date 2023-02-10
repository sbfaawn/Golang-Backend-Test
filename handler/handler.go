package handler

import (
	"Golang-Backend-Test/helper"
	"Golang-Backend-Test/logger"
	"Golang-Backend-Test/model/request"
	response2 "Golang-Backend-Test/model/response"
	"Golang-Backend-Test/service"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func HealthCheckHandler(c echo.Context) error {
	defer logger.Log("Health Check Endpoint is End", false)
	logger.Log("Health Check Endpoint is Called", false)
	response := response2.Response{Message: "Project is Up"}
	return c.JSON(http.StatusOK, response)
}

// /
func CreateUser(c echo.Context) error {
	defer logger.Log("Create User Endpoint is End", false)
	logger.Log("Create User Endpoint is Called", false)
	o := new(request.User)
	if err := c.Bind(o); err != nil {
		logger.Log("Error Occured when Create User : "+err.Error(), true)
		return err
	}

	if err := service.InsertUser(o); err != nil {
		logger.Log("Error Occured when Create User : "+err.Error(), true)
		return err
	}

	response := response2.Response{
		Message: "Success to add data",
		Data:    o,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateOrderItem(c echo.Context) error {
	defer logger.Log("Create Order Item Endpoint is End", false)
	logger.Log("Create Order Item Endpoint is Called", false)
	o := new(request.OrderItem)
	if err := c.Bind(o); err != nil {
		return err
	}

	if err := helper.IsValidDate(o.ExpiredAt); err != nil {
		logger.Log("Error Occured when Create Order Item : "+err.Error(), true)
		return err
	}

	if err := service.InsertOrderItem(o); err != nil {
		logger.Log("Error Occured when Create Order Item : "+err.Error(), true)
		return err
	}

	response := response2.Response{
		Message: "Success to add data",
		Data:    o,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateOrderHistory(c echo.Context) error {
	defer logger.Log("Create Order History Endpoint is End", false)
	logger.Log("Create Order History Endpoint is Called", false)
	o := new(request.OrderHistory)
	if err := c.Bind(o); err != nil {
		logger.Log("Error Occured when Create Order History : "+err.Error(), true)
		return err
	}

	err := service.InsertOrderHistory(o)
	if err != nil {
		logger.Log("Error Occured when Create Order History : "+err.Error(), true)
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
	defer logger.Log("List User Endpoint is End", false)
	logger.Log("List User Endpoint is Called", false)
	users, err := service.GetUsers()
	if err != nil {
		logger.Log("Error Occured when List User : "+err.Error(), true)
		return err
	}

	response := response2.Response{
		Message: "Success to Retrieve all data",
		Data:    users,
	}

	return c.JSON(http.StatusOK, response)
}

func ListOrderItems(c echo.Context) error {
	defer logger.Log("List Order Item Endpoint is End", false)
	logger.Log("List Order Item Endpoint is Called", false)
	orderItems, err := service.GetOrderItems()
	if err != nil {
		logger.Log("Error Occured when List Order History : "+err.Error(), true)
		return err
	}

	response := response2.Response{
		Message: "Success to Retrieve all data",
		Data:    orderItems,
	}

	return c.JSON(http.StatusOK, response)
}

func ListOrderHistories(c echo.Context) error {
	defer logger.Log("List Order History Endpoint is End", false)
	logger.Log("List Order Hsitory Endpoint is Called", false)
	orderHistories, err := service.GetOrderHistories()
	if err != nil {
		logger.Log("Error Occured when List Order History : "+err.Error(), true)
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
	defer logger.Log("Detail Order Item Endpoint is End", false)
	logger.Log("Detail Order Item Endpoint is Called", false)
	paramId := c.QueryParam("id")
	idInt, err := strconv.Atoi(paramId)
	if err != nil {
		logger.Log("Error Occured when Get Detail User : "+err.Error(), true)
		return err
	}

	user, err := service.GetUserById(&idInt)
	if err != nil {
		logger.Log("Error Occured when Get Detail User : "+err.Error(), true)
		return err
	}

	response := response2.Response{
		Message: "Success to Retrieve data with Id " + paramId,
		Data:    user,
	}

	return c.JSON(http.StatusOK, response)
}

func DetailOrderItem(c echo.Context) error {
	defer logger.Log("Detail Order Item Endpoint is End", false)
	logger.Log("Detail Order Item Endpoint is Called", false)
	paramId := c.QueryParam("id")
	idInt, err := strconv.Atoi(paramId)
	if err != nil {
		logger.Log("Error Occured when Get Detail Order Item : "+err.Error(), true)
		return err
	}

	orderItem, err := service.GetOrderItemById(&idInt)
	if err != nil {
		logger.Log("Error Occured when Get Detail Order Item : "+err.Error(), true)
		return err
	}

	response := response2.Response{
		Message: "Success to Retrieve data with Id " + paramId,
		Data:    orderItem,
	}

	return c.JSON(http.StatusOK, response)
}

func DetailOrderHistory(c echo.Context) error {
	defer logger.Log("Detail Order History Endpoint is End", false)
	logger.Log("Detail Order History Endpoint is Called", false)
	paramId := c.QueryParam("id")
	idInt, err := strconv.Atoi(paramId)
	if err != nil {
		logger.Log("Error Occured when Get Detail Order History : "+err.Error(), true)
		return err
	}

	orderHistory, err := service.GetOrderHistoriesById(&idInt)
	if err != nil {
		logger.Log("Error Occured when Get Detail Order History : "+err.Error(), true)
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
	defer logger.Log("Update User Endpoint is End", false)
	logger.Log("Update User Endpoint is Called", false)
	o := new(request.User)
	if err := c.Bind(o); err != nil {
		logger.Log("Error Occured when Update User : "+err.Error(), true)
		return err
	}

	if o.Id == nil {
		logger.Log("Error Occured when Update User : Id need to be specified in request", true)
		return errors.New("Id need to be specified in request")
	}

	user, err := service.UpdateUser(o)
	if err != nil {
		logger.Log("Error Occured when Update User : "+err.Error(), true)
		return err
	}

	response := response2.Response{
		Message: "Success to update data with Id " + strconv.Itoa(*o.Id),
		Data:    user,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateOrderItem(c echo.Context) error {
	defer logger.Log("Update Order Item Endpoint is End", false)
	logger.Log("Update Order Item Endpoint is Called", false)
	o := new(request.OrderItem)
	if err := c.Bind(o); err != nil {
		logger.Log("Error Occured when Update Order Item : "+err.Error(), true)
		return err
	}

	if o.Id == nil {
		logger.Log("Error Occured when Update Order Item : Id need to be specified in request", true)
		return errors.New("Id need to be specified in request")
	}

	orderItem, err := service.UpdateOrderItem(o)
	if err != nil {
		logger.Log("Error Occured when Update Order Item : "+err.Error(), true)
		return err
	}

	response := response2.Response{
		Message: "Success to update data with Id " + strconv.Itoa(*o.Id),
		Data:    orderItem,
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateOrderHistory(c echo.Context) error {
	defer logger.Log("Update Order History Endpoint is End", false)
	logger.Log("Update Order History Endpoint is Called", false)
	o := new(request.OrderHistory)
	if err := c.Bind(o); err != nil {
		logger.Log("Error Occured when Update Order History : "+err.Error(), true)
		return err
	}

	if o.Id == nil {
		logger.Log("Error Occured when Update History : Id need to be specified in request", true)
		return errors.New("Id need to be specified in request")
	}

	orderHistory, err := service.UpdateOrderHistory(o)
	if err != nil {
		logger.Log("Error Occured when Update Order History : "+err.Error(), true)
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
	defer logger.Log("Delete User Endpoint is End", false)
	logger.Log("Delete User Endpoint is Called", false)
	paramId := c.QueryParam("id")
	idInt, err := strconv.Atoi(paramId)
	if err != nil {
		logger.Log("Error Occured when Delete User : "+err.Error(), true)
		return err
	}

	err = service.SoftDeleteUser(&idInt)
	if err != nil {
		logger.Log("Error Occured when Delete User : "+err.Error(), true)
		return err
	}

	response := response2.Response{
		Message: "Success to Delete data with Id " + paramId,
		Data:    nil,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteOrderItem(c echo.Context) error {
	defer logger.Log("Delete Order Item Endpoint is End", false)
	logger.Log("Delete Order Item Endpoint is Called", false)
	paramId := c.QueryParam("id")
	idInt, err := strconv.Atoi(paramId)
	if err != nil {
		logger.Log("Error Occured when Delete Order Item : "+err.Error(), true)
		return err
	}

	err = service.SoftDeleteOrderItem(&idInt)
	if err != nil {
		logger.Log("Error Occured when Delete Order Item : "+err.Error(), true)
		return err
	}

	response := response2.Response{
		Message: "Success to Delete data with Id " + paramId,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, response)
}
