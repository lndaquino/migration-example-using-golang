package models

import "github.com/jinzhu/gorm"

type Users struct {
	ID       uint `gorm:"primaryKey"`
	UserName string
	Password string
	Email    string
}

func (u *Users) Create(db *gorm.DB) (*Users, error) {
	if err := db.Create(&u).Error; err != nil {
		return &Users{}, err
	}
	return u, nil
}
