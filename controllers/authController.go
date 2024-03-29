package controllers
import (
	"encoding/json"
	"serviceemployee/configs"
	"serviceemployee/models"
	"serviceemployee/helpers"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if register.Password != register.PasswordConfirm {
		helpers.Response(w, 400, "Password not match", nil)
		return
	}

	passwordHash, err := helpers.HashPassword(register.Password)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	employee := models.Employee {
		Name:			register.Name,
		Email:		register.Email,
		Password:	passwordHash,
	}

	if err := configs.DB.Create(&employee).Error; err!= nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 201, "Register Successfully", nil)
}

func Login(w http.ResponseWriter, r *http.Request){
	var login models.Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var employee models.Employee
	if err := configs.DB.First(&employee, "email = ?", login.Email).Error; err != nil {
		helpers.Response(w, 404, "Wrong email or password", nil)
		return
	}

	if err := helpers.VerifyPassword(employee.Password, login.Password); err != nil {
		helpers.Response(w, 404, "Wrong email or password", nil)
		return
	}

	token, err := helpers.CreateToken(&employee)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.ResponseToken(w, 200, "Successfully Login", token)

}