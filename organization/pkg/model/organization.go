package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name string
	Slug string `gorm:"unique"`
}
