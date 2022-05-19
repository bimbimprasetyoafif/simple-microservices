package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	OrgID int
	Value string
}
