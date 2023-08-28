package models

type CustomerCreateI struct {
	FullName  string `json:"fullName" validate:"required"`
	ContactNo string `json:"contactNo" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Tambon    string `json:"tambon" validate:"required"`
	Amphure   string `json:"amphure" validate:"required"`
	Province  string `json:"province" validate:"required"`
	PostCode  int    `json:"postCode" validate:"required"`
}

type CustomerCreateO struct{}

type CustomerReadI struct {
	FullName  string `json:"fullName"`
	ContactNo string `json:"contactNo"`
}

type CustomerReadO struct {
	Id        string `json:"id"`
	FullName  string `json:"fullName"`
	ContactNo string `json:"contactNo"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Tambon    string `json:"tambon"`
	Amphure   string `json:"amphure"`
	Province  string `json:"province"`
	PostCode  int    `json:"postCode"`
}

type CustomerUpdateI struct {
	Id        string `json:"id" validate:"required"`
	FullName  string `json:"fullName"`
	ContactNo string `json:"contactNo"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Tambon    string `json:"tambon"`
	Amphure   string `json:"amphure"`
	Province  string `json:"province"`
	PostCode  int    `json:"postCode"`
}

type CustomerUpdateO struct {
	Id        string `json:"id" validate:"required"`
	FullName  string `json:"fullName"`
	ContactNo string `json:"contactNo"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Tambon    string `json:"tambon"`
	Amphure   string `json:"amphure"`
	Province  string `json:"province"`
	PostCode  int    `json:"postCode"`
}

type CustomerDeleteI struct {
	Id string `json:"id" validate:"required"`
}

type CustomerDeleteO struct{}
