package service

import (
	"Golang-Backend-Test/entity"
	initializers "Golang-Backend-Test/initializer"
	"Golang-Backend-Test/model/request"
	"fmt"
	"time"
)

func InsertOrderHistory(orderHistory *request.OrderHistory) error {
	// convert to entity
	entity := entity.OrderHistory{
		UserId:       orderHistory.UserId,
		OrderItemId:  orderHistory.OrderItemId,
		Descriptions: orderHistory.Description,
	}

	// validation
	if _, err := GetUserById(&orderHistory.UserId); err != nil {
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

type OrderHistoryJoin struct {
	FullName    string    `json:"fullName"`
	FirstOrder  string    `json:"firstOrder"`
	OrderName   string    `json:"orderName"`
	Price       string    `json:"price"`
	ExpiredAt   time.Time `json:"expiredAt"`
	Description string    `json:"description"`
}

func GetOrderHistories() (OrderHistoryJoin, error) {
	var result OrderHistoryJoin

	err := initializers.DB.Table("order_histories").Select("users.full_name, users.first_order, order_items.name, order_items.price, order_items.expired_at, order_histories.description").Joins("JOIN order_items ON order_histories.order_item_id = order_items.id").Joins("JOIN users ON order_histories.user_id = users.id").Scan(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil

}
