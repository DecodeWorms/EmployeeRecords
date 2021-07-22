package types

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name   string
	Gender string
}
