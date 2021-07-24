package handlers

import (
	"employeerecords/repository"
	"employeerecords/types"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Employee struct {
}

func (emp Employee) Create(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var model types.Employee
		json.NewDecoder(r.Body).Decode(&model)

		var controller = repository.Employee{}
		controller.Create(db, model)

	}
}

func (emp Employee) GetEmployees(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var model = []types.Employee{}

		var controller = repository.Employee{}
		controller.GetEmployees(db).Scan(&model)

		json.NewEncoder(w).Encode(model)

	}
}

func (emp Employee) GetEmployee(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userParameter types.Employee
		json.NewDecoder(r.Body).Decode(&userParameter)

		var result types.Employee

		controller := repository.Employee{}
		controller.GetEmployee(db, userParameter).Scan(&result)
		json.NewEncoder(w).Encode(result)
	}
}

func (emp Employee) UpdateEmployee(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userParameter types.Employee
		json.NewDecoder(r.Body).Decode(&userParameter)

		var serverResponse types.Employee

		controller := repository.Employee{}
		controller.UpdateEmployee(db, userParameter).Scan(&serverResponse)

		json.NewEncoder(w).Encode(serverResponse)
	}
}

func (emp Employee) DeleteEmployee(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userParameter types.Employee
		json.NewDecoder(r.Body).Decode(&userParameter)

		controller := repository.Employee{}
		controller.DeleteEmployee(db, userParameter)
	}
}

func (emp Employee) CompleteInfos(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userParameter types.Employee
		json.NewDecoder(r.Body).Decode(&userParameter)

		type result struct {
			Name       string
			Gender     string
			Department string
			Salary     int
		}

		var theResult = []result{}

		controller := repository.Employee{}
		controller.CompleteInfos(db, userParameter).Scan(&theResult)
		json.NewEncoder(w).Encode(theResult)
	}
}

func (emp Employee) GetName(db *gorm.DB) http.HandlerFunc {
	var userParameter types.Employee
	json.NewDecoder(r.Body).Decode(&userParameter)

	var result types.Employee

	controller := repository.Employee{}
	controller.GetName(db, userParameter).Scan(&serverResponse)
	json.NewEncoder(w).Encode(result)
}
