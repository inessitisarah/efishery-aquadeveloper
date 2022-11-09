package repository

import (
	"mymodule/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user entity.User) error
	Delete(id int) error
	GetAll() ([]entity.User, error)
	FindById(id int) (entity.User, error)
	Update(entity.User) (entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) Create(user entity.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) Delete(id int) error {
	if err := repo.db.Debug().Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo UserRepository) FindById(id int) (entity.User, error) {
	var user entity.User
	if err := repo.db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u UserRepository) GetAll() ([]entity.User, error) {
	var users []entity.User //declare variable buat nyimpen semua data
	if err := u.db.Find(&users).Error; err != nil {
		return nil, nil
	}
	return users, nil
}

func (repo UserRepository) Update(user entity.User) (entity.User, error) {
	if err := repo.db.Debug().Save(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}
