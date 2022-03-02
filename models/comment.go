package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserRefer  int    `json:"user_id"`
	PhotoRefer int    `json:"photo_id"`
	User       User   `gorm:"foreignKey:UserRefer"`
	Photo      Photo  `gorm:"foreignKey:PhotoRefer"`
	Message    string `gorm:"not null" json:"message"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}