package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title" form:"title" valid:"required~Photo Title is Required"`
	Caption   string    `json:"caption" form:"caption"`
	PhotoURL  string    `json:"photo_url" form:"photo_url" valid:"required~Photo URL is Required"`
	UserID    int       `json:"user_id"`
	User      *User     `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	return err
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
