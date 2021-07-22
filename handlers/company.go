package handlers

import (
	"employeerecords/repository"
	"employeerecords/types"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Company struct{}

func (comp Company) Create(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var userParameter types.Company
		json.NewDecoder(r.Body).Decode(&userParameter)

		controller := repository.Company{}
		controller.Create(db, userParameter)

	}
}

func (com Company) GetEmployee(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userParameter types.Company
		json.NewDecoder(r.Body).Decode(&userParameter)

		var serverResponse types.Company
		controller := repository.Company{}
		controller.GetEmployee(db, userParameter).Scan(&serverResponse)
		json.NewEncoder(w).Encode(serverResponse)
	}
}

func (com Company) GetEmployees(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var model = []types.Company{}

		controller := repository.Company{}
		controller.GetEmployees(db).Scan(&model)
		json.NewEncoder(w).Encode(model)
	}
}

func (com Company) UpdateRecord(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userParameter types.Company
		json.NewDecoder(r.Body).Decode(&userParameter)

		var serverResponse types.Company

		controller := repository.Company{}
		controller.UpdateRecord(db, userParameter).Scan(&serverResponse)

		json.NewEncoder(w).Encode(serverResponse)
	}
}

func (comp Company) DeleteRecord(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userParameter types.Company
		json.NewDecoder(r.Body).Decode(&userParameter)

		var serverResponse types.Company

		controller := repository.Company{}
		controller.DeleteRecord(db, userParameter).Scan(&serverResponse)

		json.NewEncoder(w).Encode(serverResponse)
	}
}

func (com Company) EmployeeInfos(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userParameter types.Company
		json.NewDecoder(r.Body).Decode(&userParameter)

		type result struct {
			Name       string
			Gender     string
			Department string
			Salary     int
		}

		var serverResponse = []result{}

		controller := repository.Company{}
		thevalue := controller.EmployeeInfos(db, userParameter).Scan(&serverResponse)
		fmt.Print(thevalue)

	}
}
