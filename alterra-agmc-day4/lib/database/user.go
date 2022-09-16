package database

import (
	"github.com/danielsugianto/alterra-agmc-day4/config"
	"github.com/danielsugianto/alterra-agmc-day4/middlewares"
	"github.com/danielsugianto/alterra-agmc-day4/models"
)

// get all users
func GetUsers() ([]models.User, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// create new user
func CreateUser(user models.User) (models.User, error) {

	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// get user by ID
func GetUser(id string) (models.User, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// delete user by ID
func DeleteUser(id string) error {
	var user models.User

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

// update user by ID
func UpdateUser(id string, userParam models.User) (models.User, error) {
	var user models.User
	user, err := GetUser(id)
	if err != nil {
		return user, err
	}
	user.Email = userParam.Email
	user.Password = userParam.Password
	user.Name = userParam.Name
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//Login User
func LoginUser(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password= ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	return user, nil
}
