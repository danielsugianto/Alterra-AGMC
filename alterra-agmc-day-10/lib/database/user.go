package database

import (
	"github.com/danielsugianto/alterra-agmc-day-10/middlewares"
	"github.com/danielsugianto/alterra-agmc-day-10/models"
	"gorm.io/gorm"
)

type mysqlUsersRepository struct {
	DB *gorm.DB
}

// NewMySQLUsersRepository will create an object that represent the users.Repository interface
func NewMySQLUsersRepository(Conn *gorm.DB) models.UsersMySQLRepository {
	return &mysqlUsersRepository{Conn}
}

// get all users
func (mysqlUsersRepository *mysqlUsersRepository) GetUsers() ([]models.User, error) {
	var users []models.User

	if err := mysqlUsersRepository.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// create new user
func (mysqlUsersRepository *mysqlUsersRepository) CreateUser(user models.User) (models.User, error) {

	if err := mysqlUsersRepository.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// get user by ID
func (mysqlUsersRepository *mysqlUsersRepository) GetUser(id string) (models.User, error) {
	var user models.User

	if err := mysqlUsersRepository.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// delete user by ID
func (mysqlUsersRepository *mysqlUsersRepository) DeleteUser(id string) error {
	var user models.User

	if err := mysqlUsersRepository.DB.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

// update user by ID
func (mysqlUsersRepository *mysqlUsersRepository) UpdateUser(id string, userParam models.User) (models.User, error) {
	var user models.User
	user, err := mysqlUsersRepository.GetUser(id)
	if err != nil {
		return user, err
	}
	user.Email = userParam.Email
	user.Password = userParam.Password
	user.Name = userParam.Name
	if err := mysqlUsersRepository.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//Login User
func (mysqlUsersRepository *mysqlUsersRepository) LoginUser(user *models.User) (interface{}, error) {
	var err error
	if err = mysqlUsersRepository.DB.Where("email = ? AND password= ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	return user, nil
}
