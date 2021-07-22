package main

import (
	"employeerecords/driver"
	"employeerecords/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	db := driver.ConnDb()

	controller := handlers.Employee{}
	controller2 := handlers.Company{}

	router := mux.NewRouter()

	router.HandleFunc("/create", controller.Create(db)).Methods("POST")
	router.HandleFunc("/employees", controller.GetEmployees(db)).Methods("GET")
	router.HandleFunc("/employee", controller.GetEmployee(db)).Methods("GET")
	router.HandleFunc("/update", controller.UpdateEmployee(db)).Methods("PUT")
	router.HandleFunc("/drop", controller.DeleteEmployee(db)).Methods("DELETE")

	router.HandleFunc("/companysignup", controller2.Create(db)).Methods("POST")
	router.HandleFunc("/getemployee", controller2.GetEmployee(db)).Methods("GET")
	router.HandleFunc("/getemployees", controller2.GetEmployees(db)).Methods("GET")
	router.HandleFunc("/updateemployee", controller2.UpdateRecord(db)).Methods("PUT")
	router.HandleFunc("/deleteemployee", controller2.DeleteRecord(db)).Methods("DELETE")
	router.HandleFunc("/infos", controller2.EmployeeInfos(db)).Methods("GET")
	http.ListenAndServe(":8000", router)

}
