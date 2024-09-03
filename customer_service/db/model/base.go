package model

type Customer struct {
	BaseModel
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Rating int `json:"rating"`
	Trips []Trip `json:"trips"`
}

type Trip struct {
	BaseModel
	Source string
	Destination string 
	Cost float64 
	PaymentMethod string
	Status int	
	// add custoemr id here with foerign key so we can join with custoemr table and get all the trips?
}
