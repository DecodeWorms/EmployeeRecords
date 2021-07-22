package types

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Department string
	Salary     int
	EmployeeId int
}
