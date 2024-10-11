package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `gorm:"size:255"`
	Email          string `gorm:"size:255;unique;index;not null"`
	HashedPassword string `gorm:"size:255; not null"`
}
