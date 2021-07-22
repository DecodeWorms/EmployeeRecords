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

func (comp Company) EmployeeInfos(db *gorm.DB, model types.Company) *gorm.DB {

	result := db.Model(&types.Company{}).Select("companies.department, companies.salary, employees.name,employees.gender").Joins("left join employees on companies.employee_id = employees.id").Where("companies.employee_id = ?", model.EmployeeId)
	return result

}

// func (control Controller) Getuserandcard(db *gorm.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var theusers usermodel.User
// 		json.NewDecoder(r.Body).Decode(&theusers)

// 		type result struct {
// 			Name       string
// 			Gender     string
// 			CardType   string
// 			CardNumber string
// 		}

// 		var theresults = []result{}

// 		db.Model(&usermodel.User{}).Select("users.name,users.gender,atmcards.card_type,atmcards.card_number").Joins("left join atmcards on users.id = atmcards.user_id").Where("users.id = ?", theusers.ID).Scan(&theresults)
// 		json.NewEncoder(w).Encode(theresults)
// 	}
// }
