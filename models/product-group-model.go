package models

type CreateProductGroupI struct {
	GroupName string `json:"groupName" validate:"required"`
}

type CreateProductGroupO struct{}

type ReadProductGroupI struct {
	GroupName string `json:"groupName" validate:"required"`
}

type ReadProductGroupO struct {
	Id        string `json:"id"`
	GroupName string `json:"groupName"`
}

type UpdateProductGroupI struct {
	Id        string `json:"id" validate:"required"`
	GroupName string `json:"groupName"`
}

type UpdateProductGroupO struct {
	Id        string `json:"id"`
	GroupName string `json:"groupName"`
}

type DeleteProductGroupI struct {
	Id string `json:"id"`
}

type DeleteProductGroupO struct{}
