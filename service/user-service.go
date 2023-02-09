package service

import (
	"Golang-Backend-Test/entity"
	initializers "Golang-Backend-Test/initializer"
	"Golang-Backend-Test/model/request"
	"fmt"
)

func InsertUser(user *request.User) error {
	// convert to entity
	entity := entity.User{
		FullName:   user.FullName,
		FirstOrder: user.FirstOrder,
		DeletedAt:  nil,
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

func GetUserById(userId *int) error {
	var user entity.User
	err := initializers.DB.Where("deleted_at IS NULL").First(&user, userId).Error

	if err != nil {
		return err
	}
	fmt.Println(user)

	return nil
}
