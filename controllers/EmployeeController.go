package controllers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"serviceemployee/configs"
	"serviceemployee/models"
	"serviceemployee/helpers"
	"github.com/gorilla/mux"
)

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	employee := r.Context().Value("employeeinfo").(*helpers.MyCustomClaims)
	employeeResponse := &models.MyProfile {
		ID:			employee.ID,
		Name:		employee.Name,
		Email:	employee.Email,
	}
	helpers.Response(w, 200, "My Profile", employeeResponse)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var register models.Register
	vars := mux.Vars(r)
	employeeId := vars["id"]

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()
	
	fmt.Println(employeeId)

	if register.Name == "" || register.Email == "" {
		helpers.Response(w, http.StatusBadRequest, "Name and Email are required fields", nil)
		return
	}

	employee := models.Employee {
		Name:			register.Name,
		Email:		register.Email,
	}

	if err := configs.DB.First(&employee, "id = ?", employeeId).Error; err != nil {
		helpers.Response(w, 404, "Employee not found", nil)
		return
	}

	// Update employee information
	if err := configs.DB.Model(&models.Employee{}).Where("id = ?", employeeId).Updates(&employee).Error; err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Update Successfully", nil)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request){
	var employee models.Employee

	vars := mux.Vars(r)
	employeeId := vars["id"]

	if err := configs.DB.First(&employee, "id = ?", employeeId).Error; err != nil {
		helpers.Response(w, 404, "Employee not found", nil)
		return
	}

	if err := configs.DB.Delete(&employee, employeeId).Error; err != nil {
		helpers.Response(w, 404, "Employee not found", nil)
		return
	}

	helpers.Response(w, 200, "Delete Successfully", nil)
}