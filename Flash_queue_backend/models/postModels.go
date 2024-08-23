package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string
	Number   string
	Position int
}
