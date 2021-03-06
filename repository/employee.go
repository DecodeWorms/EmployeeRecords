package repository

import (
	"employeerecords/types"

	"gorm.io/gorm"
)

type Employee struct {
}

func (emp Employee) Create(db *gorm.DB, model types.Employee) {

	db.Create(&model)

}

func (emp Employee) GetEmployees(db *gorm.DB) *gorm.DB {

	var model = []types.Employee{}
	result := db.Find(&model)
	return result
}

func (emp Employee) GetEmployee(db *gorm.DB, model types.Employee) *gorm.DB {

	result := db.Where(" name = ?", model.Name).Find(&model)
	return result
}

func (emp Employee) UpdateEmployee(db *gorm.DB, model types.Employee) *gorm.DB {

	result := db.Model(&types.Employee{}).Where("gender = ? ", model.Gender).Update("gender", "male")

	return result
}

func (emp Employee) DeleteEmployee(db *gorm.DB, model types.Employee) *gorm.DB {

	result := db.Where("name = ?", model.Name).Delete(&model)
	return result

}

func (emp Employee) CompleteInfos(db *gorm.DB, model types.Employee) *gorm.DB {

	result := db.Model(&types.Employee{}).Select("employees.name, employees.gender, companies.department, companies.salary").Joins("left join companies on employees.id = companies.employee_id").Where("employees.id = ?", model.ID)
	return result
}

func (emp Employee) GetName(db *gorm.DB, model types.Employee) *gorm.DB {
	return db.Where("name = ?", model.Name).Find(&model)
}
