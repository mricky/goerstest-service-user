package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user Users) (Users, error)
	FindByEmail(name string, password string) (Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db} // return struct repository yang sudah ada object db
}

func (r *repository) Save(user Users) (Users, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
}

func (r *repository) FindByEmail(email string, password string) (Users, error) {
	var user Users

	err := r.db.Where("email = ? AND password_hash = ?", email, password).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}
