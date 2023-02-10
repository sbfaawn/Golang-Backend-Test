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

func GetUserById(userId *int) (entity.User, error) {
	result, err2 := initializers.Client.Get(initializers.Ctx, "user_"+strconv.Itoa(*userId)).Result()
	if err2 == nil {
		fmt.Println("Get From Cache")
		user := entity.User{}
		err := json.Unmarshal([]byte(result), &user)
		if err != nil {
			return user, err
		}

		return user, nil
	}

	var user entity.User
	err := initializers.DB.Where("deleted_at IS NULL").First(&user, userId).Error

	if err != nil {
		return user, err
	}
	fmt.Println(user)
	marshal, err2 := json.Marshal(user)
	if err2 != nil {
		return user, err2
	}

	err2 = initializers.Client.Set(initializers.Ctx, "user_"+strconv.Itoa(*userId), string(marshal), 30*time.Minute).Err()
	if err2 != nil {
		return user, err2
	}

	return user, nil
}

func GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := initializers.DB.Where("deleted_at IS NULL").Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}

func SoftDeleteUser(id *int) error {
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

func UpdateUser(user *request.User) (entity.User, error) {
	entity := entity.User{
		Id: *user.Id,
	}

	err := initializers.DB.Model(&entity).Updates(map[string]interface{}{
		"full_name":   user.FullName,
		"first_order": user.FirstOrder,
	}).Error
	if err != nil {
		return entity, err
	}

	return entity, nil
}
