package service

import (
	"Golang-Backend-Test/entity"
	"Golang-Backend-Test/helper"
	initializers "Golang-Backend-Test/initializer"
	"Golang-Backend-Test/model/request"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
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
	result, err2 := initializers.Client.Get(initializers.Ctx, "orderItem_"+strconv.Itoa(*orderItemId)).Result()
	if err2 == nil {
		fmt.Println("Get From Cache")
		orderItem := entity.OrderItem{}
		err := json.Unmarshal([]byte(result), &orderItem)
		if err != nil {
			return orderItem, err
		}

		return orderItem, nil
	}

	var orderItem entity.OrderItem
	err := initializers.DB.Where("deleted_at IS NULL").First(&orderItem, orderItemId).Error

	if err != nil {
		return orderItem, err
	}
	fmt.Println(orderItem)
	marshal, err2 := json.Marshal(orderItem)
	if err2 != nil {
		return orderItem, err2
	}

	err2 = initializers.Client.Set(initializers.Ctx, "orderItem_"+strconv.Itoa(*orderItemId), string(marshal), 30*time.Minute).Err()
	if err2 != nil {
		return orderItem, err2
	}

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

func SoftDeleteOrderItem(id *int) error {
	user := entity.User{
		Id: *id,
	}

	err := initializers.DB.Model(&user).Updates(map[string]interface{}{
		"deleted_at": time.Now(),
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateOrderItem(orderItem *request.OrderItem) (entity.OrderItem, error) {
	entity := entity.OrderItem{
		Id: *orderItem.Id,
	}

	err := initializers.DB.Model(&entity).Updates(map[string]interface{}{
		"name":       orderItem.Name,
		"price":      orderItem.Price,
		"expired_at": orderItem.ExpiredAt,
	}).Error
	if err != nil {
		return entity, err
	}

	return entity, nil
}
