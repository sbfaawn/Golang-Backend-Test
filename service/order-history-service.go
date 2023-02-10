package service

import (
	"Golang-Backend-Test/entity"
	initializers "Golang-Backend-Test/initializer"
	"Golang-Backend-Test/model/request"
	"encoding/json"
	"fmt"
	"strconv"
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

	if _, err := GetOrderItemById(&orderHistory.OrderItemId); err != nil {
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
	OrderName   string    `json:"orderName" gorm:"column:name"`
	Price       string    `json:"price"`
	ExpiredAt   time.Time `json:"expiredAt"`
	Description string    `json:"description"`
}

func GetOrderHistories() (OrderHistoryJoin, error) {
	var result OrderHistoryJoin

	err := initializers.DB.Table("order_histories").Where("users.deleted_at IS NULL").Where("order_items.deleted_at IS NULL").Select("users.full_name, users.first_order, order_items.name, order_items.price, order_items.expired_at, order_histories.description").Joins("JOIN order_items ON order_histories.order_item_id = order_items.id").Joins("JOIN users ON order_histories.user_id = users.id").Scan(&result).Error
	fmt.Println(result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func GetOrderHistoriesById(id *int) (OrderHistoryJoin, error) {
	cache, err2 := initializers.Client.Get(initializers.Ctx, "orderHistory_"+strconv.Itoa(*id)).Result()
	if err2 == nil {
		fmt.Println("Get From Cache")
		orderHistoryJoin := OrderHistoryJoin{}
		err := json.Unmarshal([]byte(cache), &orderHistoryJoin)
		if err != nil {
			return orderHistoryJoin, err
		}

		return orderHistoryJoin, nil
	}

	var result OrderHistoryJoin

	err := initializers.DB.Table("order_histories").Where("order_histories.id = ?", id).Where("users.deleted_at IS NULL").Where("order_items.deleted_at IS NULL").Select("users.full_name, users.first_order, order_items.name, order_items.price, order_items.expired_at, order_histories.description").Joins("JOIN order_items ON order_histories.order_item_id = order_items.id").Joins("JOIN users ON order_histories.user_id = users.id").Scan(&result).Error
	if err != nil {
		return result, err
	}

	fmt.Println(result)
	marshal, err2 := json.Marshal(result)
	if err2 != nil {
		return result, err2
	}

	err2 = initializers.Client.Set(initializers.Ctx, "orderHistory_"+strconv.Itoa(*id), string(marshal), 30*time.Minute).Err()
	if err2 != nil {
		return result, err2
	}

	return result, nil
}

func UpdateOrderHistory(history *request.OrderHistory) (entity.OrderHistory, error) {
	entity := entity.OrderHistory{
		Id: *history.Id,
	}

	// validation
	if _, err := GetUserById(&history.UserId); err != nil {
		return entity, err
	}

	if _, err := GetOrderItemById(&history.OrderItemId); err != nil {
		return entity, err
	}

	err := initializers.DB.Model(&entity).Updates(map[string]interface{}{
		"user_id":       history.UserId,
		"order_item_id": history.OrderItemId,
		"description":   history.Description,
	}).Error
	if err != nil {
		return entity, err
	}

	return entity, nil
}
