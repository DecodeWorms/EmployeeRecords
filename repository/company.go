package repository

import (
	"employeerecords/types"

	"gorm.io/gorm"
)

type Company struct{}

func (comp Company) Create(db *gorm.DB, model types.Company) {

	db.Create(&model)

}

func (com Company) GetEmployee(db *gorm.DB, model types.Company) *gorm.DB {
	result := db.Where("employee_id = ?", model.EmployeeId).Find(&model)
	return result
}

func (com Company) GetEmployees(db *gorm.DB) *gorm.DB {
	var model = []types.Company{}

	result := db.Find(&model)
	return result

}

func (comp Company) UpdateRecord(db *gorm.DB, model types.Company) *gorm.DB {

	result := db.Model(&types.Company{}).Where("employee_id = ?", model.EmployeeId).Update("salary", 200000)
	return result
}

func (com Company) DeleteRecord(db *gorm.DB, model types.Company) *gorm.DB {
	result := db.Where("employee_id = ?", model.EmployeeId).Delete(&model)
	return result
}
