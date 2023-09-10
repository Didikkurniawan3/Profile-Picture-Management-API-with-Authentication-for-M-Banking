package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"not null" json:"username" form:"username" valid:"required~Username is required"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid Email"`
	Password  string    `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password minimum is 6 characters"`
	Photo     []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photo,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	return err
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	return err
}
