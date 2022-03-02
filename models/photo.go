package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title     string `gorm:"not null" json:"title" valid:"required~Title is required"`
	Caption   string `json:"caption"`
	PhotoURL  string `gorm:"not null" json:"photo_url" valid:"required~Photo URL is required"`
	UserRefer int    `json:"user_id"`
	User      User   `gorm:"foreignKey:UserRefer"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}