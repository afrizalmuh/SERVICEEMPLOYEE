package controllers

import (
	"encoding/json"
	"serviceemployee/configs"
	"net/http"
	"serviceemployee/models"
	"serviceemployee/helpers"
	"github.com/gorilla/mux"
	"fmt"
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

	employee := models.Employee {
		Name:			register.Name,
		Email:		register.Email,
	}

	if err := configs.DB.Where("id", employeeId).Updates(&employee).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil) 
		return
	}

	helpers.Response(w, 201, "Update Successfully", nil)
}
