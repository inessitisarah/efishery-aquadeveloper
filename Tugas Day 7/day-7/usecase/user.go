package usecase

import (
	"mymodule/entity"
	"mymodule/entity/response"
	"mymodule/repository"

	"github.com/jinzhu/copier"
)

type IUserUseCase interface {
	CreateUser(user response.CreateUserRequest) error
	GetListUser() ([]response.GetUserResponse, error)
	DeleteUser(id int) error
	UpdateUser(id int) error
}

type UserUseCase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepository: userRepository}
}

/* func (u UserUseCase) CreateUser(req response.CreateUserRequest) error {
	user := entity.User{}
	copier.Copy(&user, &req)
} */

func (u UserUseCase) CreateUser(req response.CreateUserRequest) error {
	users := entity.User{}

	copier.Copy(&users, &req)
	err := u.userRepository.Create(users)
	if err != nil {
		return err
	}
	return nil
}

// msh bingung kl dideleteitu diapain krn gamungkin dicopy
func (u UserUseCase) DeleteUser(id int) error {
	_, err := u.userRepository.FindById(id)
	if err != nil {
		return err
	}

	err = u.userRepository.Delete(id)
	return err
}

func (services UserUseCase) UpdateUser(userRequest entity.UpdateUserRequest, id int) (response.UserResponse, error) {
	user, err := services.userRepository.FindById(id)

	if err != nil {
		return response.UserResponse{}, err
	}

	copier.CopyWithOption(&user, &userRequest, copier.Option{IgnoreEmpty: true})

	user, err = services.userRepository.Update(user)

	userRes := response.UserResponse{}

	copier.Copy(&userRes, &user)

	return userRes, nil
}

func (u UserUseCase) GetList() ([]response.GetUserResponse, error) {
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	userResponse := []response.GetUserResponse{}
	copier.Copy(&userResponse, &users)
	return userResponse, nil
}
