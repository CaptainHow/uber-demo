package dto

import "github.com/uber-demo/customer/db/model"

type Customer struct {
	model.BaseModel
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Rating int `json:"rating"`
}

type RegisterCutomer struct {
	model.BaseModel
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Rating int `json:"rating"`
}