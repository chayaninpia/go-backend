package models

type BaseProvinceModel struct {
	NameTh string `json:"name_th"`
	NameEn string `json:"name_en"`
}

type BaseAmphureModel struct {
	NameTh   string            `json:"name_th"`
	NameEn   string            `json:"name_en"`
	Province BaseProvinceModel `json:"province"`
}

type AddressCreateI struct {
	PostCode int              `json:"post_code"`
	NameTh   string           `json:"name_th"`
	NameEn   string           `json:"name_en"`
	Amphure  BaseAmphureModel `json:"amphure"`
}

type AddressCreateO struct{}

type AddressReadI struct {
	Filter string `json:"filter"`
}

type AddressReadO struct {
	Id       string           `json:"id"`
	PostCode int              `json:"post_code"`
	NameTh   string           `json:"name_th"`
	NameEn   string           `json:"name_en"`
	Amphure  BaseAmphureModel `json:"amphure"`
}

type AddressUpdateI struct {
	Id       string           `json:"id"`
	PostCode int              `json:"post_code"`
	NameTh   string           `json:"name_th"`
	NameEn   string           `json:"name_en"`
	Amphure  BaseAmphureModel `json:"amphure"`
}

type AddressUpdateO struct {
	Id       string           `json:"id"`
	PostCode int              `json:"post_code"`
	NameTh   string           `json:"name_th"`
	NameEn   string           `json:"name_en"`
	Amphure  BaseAmphureModel `json:"amphure"`
}

type AddressDeleteI struct {
	Id string `json:"id" validate:"required"`
}

type AddressDeleteO struct{}
