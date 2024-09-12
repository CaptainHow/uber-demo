package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/uber-demo/customer/api/v1/dto"
	"github.com/uber-demo/customer/api/v1/helper"
	database "github.com/uber-demo/customer/db"
	"gorm.io/gorm"
)


func GetCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	// fmt.Println(id, len(customersArray))
	// if id >= len(customersArray) || id < 0 {
	// 	respondWithError(w, http.StatusBadRequest, "User does not exists")
	// 	return
	// }
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"msg":"Customer " + strconv.Itoa(id)})
	// OR
	// render.JSON(w, r, customer) 	
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer dto.RegisterCutomer

	err := render.DecodeJSON(r.Body, &customer)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	db := database.GetDb()
	// Check if User already exists
	start := time.Now()
	logger := log.Default()
	logger.Println(start)
	err = db.Transaction(func(tx *gorm.DB) error {
		var user string
		tx.Model(&dto.Customer{}).Where("username=?", customer.Username).Select("username").Find(&user)
		fmt.Println(user)
		if user != "" {
			return errors.New("username already exists")
		}
		return nil
	})
	
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return 
	}

	customer.Password, _ = helper.HashPassword(customer.Password)

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&dto.Customer{}).Create(&customer).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	// var custID = customer.Save()
	helper.RespondWithJSON(w, http.StatusOK, map[string]uuid.UUID{"CustID" : customer.Id})
}