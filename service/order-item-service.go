package service

import (
	"Golang-Backend-Test/entity"
	"Golang-Backend-Test/helper"
	initializers "Golang-Backend-Test/initializer"
	"Golang-Backend-Test/model/request"
	"fmt"
)

func InsertOrderItem(order *request.OrderItem) error {
	// convert to entity
	entity := entity.OrderItem{
		Name:      order.Name,
		Price:     order.Price,
		ExpiredAt: helper.ConvertToTimeObject(order.ExpiredAt),
	}
	// call db & insert
	err := initializers.DB.Create(&entity).Error
	if err != nil {
		return err
	}
	fmt.Println(entity)

	/*
		json, err := json.Marshal(entity)
		if err != nil {
			return err
		}

		err = initializers.Client.Set(initializers.Ctx, "user_"+strconv.Itoa(entity.Id), json, 20*time.Minute).Err()
		if err != nil {
			return err
		}
	*/

	return nil
}

func GetOrderItemById(orderItemId *int) (entity.OrderItem, error) {
	var orderItem entity.OrderItem
	err := initializers.DB.Where("deleted_at IS NULL").First(&orderItem, orderItemId).Error

	if err != nil {
		return orderItem, err
	}
	fmt.Println(orderItem)

	return orderItem, nil
}

func GetOrderItems() ([]entity.OrderItem, error) {
	var orderItems []entity.OrderItem
	err := initializers.DB.Where("deleted_at IS NULL").Find(&orderItems).Error

	if err != nil {
		return orderItems, err
	}

	return orderItems, nil
}
