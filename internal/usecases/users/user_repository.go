package users

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]User, error)
	CreateUser(user User) error
	GetUserByUsername(username string) (*User, error)
	DeleteUser(username string) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) GetUsers() ([]User, error) {
	var users []User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepositoryImpl) CreateUser(user *User) error {
	err := u.db.Omit("ID").Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImpl) GetUserByUsername(username string) (*User, error) {
	var user User
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepositoryImpl) DeleteUser(username string) error {
	err := u.db.Where("username = ?", username).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}
