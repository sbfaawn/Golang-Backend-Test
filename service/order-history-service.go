package service

import (
	"Golang-Backend-Test/entity"
	initializers "Golang-Backend-Test/initializer"
	"Golang-Backend-Test/model/request"
	"fmt"
)

func InsertOrderHistory(orderHistory *request.OrderHistory) error {
	// convert to entity
	entity := entity.OrderHistory{
		UserId:       orderHistory.UserId,
		OrderItemId:  orderHistory.OrderItemId,
		Descriptions: orderHistory.Description,
	}

	// validation
	if err := GetUserById(&orderHistory.UserId); err != nil {
		return err
	}

	if err := GetOrderItemById(&orderHistory.OrderItemId); err != nil {
		return err
	}
	// call db & insert
	err := initializers.DB.Create(&entity).Error
	if err != nil {
		return err
	}
	fmt.Println(entity)

	return nil
}
