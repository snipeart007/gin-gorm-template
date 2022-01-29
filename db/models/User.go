package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	FirstName string `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName  string `json:"last_name" gorm:"type:varchar(100);not null"`
	Email     string `json:"email" gorm:"type:varchar(100);not null;unique"`
	Password  string `json:"password" gorm:"type:varchar(100);not null"`
}
