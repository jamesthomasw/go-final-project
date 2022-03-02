package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null,uniqueIndex" json:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null,uniqueIndex" json:"email" valid:"required~Your email is required"`
	Password string `gorm:"not null" json:"-" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int    `gorm:"not null" json:"age" valid:"required~Your age is required~range(8|100)~Age has to be at least 8 or higher"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
