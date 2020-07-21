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
	Name     string `gorm:"type:varchar(48);not null;unique_index"`
	Email    string `gorm:"type:varchar(48);not null;unique_index"`
	Password string `gorm:"type:varchar(128);not null"`
	Avatar   string `gorm:"type:varchar(256)"`
}

func AddMenger(name, email, password, avatar string) error {
	menger := Menger{
		Name:     name,
		Email:    email,
		Password: password,
		Avatar:   avatar,
	}
	if err := db.Create(&menger).Error; err != nil {
		return err
	}
	return nil
}

func GetMengerByEmailOrName(name, email string) (*Menger, error) {
	var menger Menger
	err := db.Where("name = ? or email = ?", name, email).Where("deleted_at is null").First(&menger).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &menger, nil
}

func GetMengerByEmail(email string) (*Menger, error) {
	var menger Menger
	err := db.Where("email = ?", email).First(&menger).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &menger, nil
}

func GetMengerByName(name string) (*Menger, error) {
	var menger Menger
	err := db.Where("name = ?", name).First(&menger).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &menger, nil
}

func GetMengerByIds(mengerIds []int64) ([]*Menger, error) {
	var mengers []*Menger
	err := db.Where("id in (?) and deleted_at is null", mengerIds).First(&mengers).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return mengers, nil
}

func GetMengerById(mengerId int64) (*Menger, error) {
	var menger Menger
	err := db.Where("id = ? and deleted_at is null", mengerId).First(&menger).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &menger, nil
}
