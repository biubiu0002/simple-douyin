package repository

import (
	"gorm.io/gorm"
)

type User struct {
	Id       int64  `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func AddUser(user *User) (*User, error) {
	err := db.Create(user).Error
	return user, err
}

func SelectByName(username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func SelectById(id int64) (*User, error) {
	var user User
	err := db.Where("id = ?", id).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
