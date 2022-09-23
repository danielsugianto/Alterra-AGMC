package usecase

import (
	"github.com/danielsugianto/alterra-agmc-day-10/models"
)

type usersUsecase struct {
	usersMySQLRepo models.UsersMySQLRepository
}

// NewUsersUsecase will create new an usersUsecase object representation of domain.UsersUsecase interface
func NewUsersUsecase(usersMySQLRepo models.UsersMySQLRepository) models.UsersUsecase {
	return &usersUsecase{
		usersMySQLRepo: usersMySQLRepo,
	}
}

// get all users
func (usersUsecase *usersUsecase) GetUsersUsecase() ([]models.User, error) {
	var users []models.User

	users, err := usersUsecase.usersMySQLRepo.GetUsers()
	if err != nil {
		return users, err
	}
	return users, nil
}

// create new user
func (usersUsecase *usersUsecase) CreateUserUsecase(userParam models.User) (models.User, error) {
	user, err := usersUsecase.usersMySQLRepo.CreateUser(userParam)
	if err != nil {
		return user, err
	}
	return user, nil
}

// get user by ID
func (usersUsecase *usersUsecase) GetUserUsecase(id string) (models.User, error) {
	user, err := usersUsecase.usersMySQLRepo.GetUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// delete user by ID
func (usersUsecase *usersUsecase) DeleteUserUsecase(id string) error {
	errDelete := usersUsecase.usersMySQLRepo.DeleteUser(id)
	if errDelete != nil {
		return errDelete
	}
	return nil
}

// update user by ID
func (usersUsecase *usersUsecase) UpdateUserUsecase(id string, userParam models.User) (models.User, error) {
	user, errUpdate := usersUsecase.usersMySQLRepo.UpdateUser(id, userParam)
	if errUpdate != nil {
		return user, errUpdate
	}
	return user, nil
}

// Login Usecase
func (usersUsecase *usersUsecase) LoginUsersUsecase(user models.User) (interface{}, error) {
	users, e := usersUsecase.usersMySQLRepo.LoginUser(&user)
	if e != nil {
		return users, e
	}
	return users, nil
}
