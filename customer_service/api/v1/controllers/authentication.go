package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/uber-demo/customer/api/v1/dto"
	"github.com/uber-demo/customer/api/v1/helper"
	database "github.com/uber-demo/customer/db"
	"gorm.io/gorm"
)

func LoginCustomer(w http.ResponseWriter, r *http.Request) {
	var customer dto.AuthenticateCustomer
	err := render.DecodeJSON(r.Body, &customer)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	db := database.GetDb()
	var customerFound dto.Customer
	err = db.Transaction(func(tx *gorm.DB) error {
		tx.Model(&dto.Customer{}).Where("username=?", customer.Username).Find(&customerFound)
		return nil
	})

	
	// User doesn't exists
	if customerFound.Username == "" || err != nil{
		helper.RespondWithError(w, http.StatusUnauthorized, "Invalid Username or Password " + err.Error())
		return
	}
	
	// Invalid Credentials
	if !helper.VerifyPassword(customer.Password, customerFound.Password) {
		helper.RespondWithError(w, http.StatusUnauthorized, "Invalid Username or Password")
		return
	}

	token, err := helper.GenerateJWTToken(customerFound.Id)
	fmt.Println(token)
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, map[string]uuid.UUID{"Id": customerFound.Id})
}