package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Menger struct {
	Model
	Name     string `gorm:"type:varchar(48);not null"`
	Email    string `gorm:"type:varchar(48);not null"`
	Password string `gorm:"type:varchar(128);not null"`
	Salt     string `gorm:"type:char(4);not null"`
	Avatar   string  `gorm:"type:varchar(256)"`
}

func AddMenger(name, email, password, salt, avatar string) error {
	menger := Menger{
		Name:     name,
		Email:    email,
		Password: password,
		Salt:     salt,
		Avatar:   avatar,
	}
	if err := db.Create(&menger).Error; err != nil {
		return err
	}
	return nil
}

func GetMengerByEmailOrName(name, email string) (*Menger, error) {
	var menger Menger
	err := db.Where("name = ? or email = ?", name, email).Where("delete_at = ?", 0).First(&menger).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &menger, nil
}

func GetMengerByEmail(email string) (*Menger, error) {
	var menger Menger
	err := db.Where("email = ? and delete_at = ?", email, 0).First(&menger).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &menger, nil
}

func GetMengerByName(name string) (*Menger, error) {
	var menger Menger
	err := db.Where("name = ? and delete_at = ?", name, 0).First(&menger).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &menger, nil
}

func GetMengerByIds(mengerIds []int64) ([]*Menger, error) {
	var mengers []*Menger
	err := db.Where("id in (?)", mengerIds).First(&mengers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return mengers, nil
}
