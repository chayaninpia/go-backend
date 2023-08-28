package models

type CreateProductI struct {
	ProductName   string                     `json:"productName"  validate:"required"`
	ProductNameTh string                     `json:"productNameTh"  validate:"required"`
	BarcodeId     string                     `json:"barcodeId"  validate:"required"`
	GroupId       string                     `json:"groupId"  validate:"required"`
	BasePrice     float64                    `json:"basePrice"  validate:"required"`
	SalePrice     float64                    `json:"salePrice"  validate:"required"`
	Unit          string                     `json:"unit"  validate:"required"`
	IsBaseProduct bool                       `json:"isBaseProduct" `
	SubProduct    []CreateProductISubProduct `json:"subProduct"`
}

type CreateProductISubProduct struct {
	BaseProductId  string `json:"baseProductId"`
	BaseProductUse int    `json:"baseProductUse"`
}

type CreateProductO struct{}

type ReadProductI struct {
	ProductName string `json:"productName" validate:"required"`
	ProductId   string `json:"productId" `
}

type ReadProductO struct {
	Id            string            `json:"id" `
	ProductName   string            `json:"productName" `
	ProductNameTh string            `json:"productNameTh" `
	BarcodeId     string            `json:"barcodeId" `
	GroupId       string            `json:"groupId" `
	BasePrice     float64           `json:"basePrice" `
	SalePrice     float64           `json:"salePrice" `
	Unit          string            `json:"unit" `
	IsBaseProduct bool              `json:"isBaseProduct" `
	ProductSub    []ReadProductOSub `json:"productSub"`
}

type ReadProductOSub struct {
	Id             string `json:"id" `
	ProductId      string `json:"productId" `
	BaseProductId  string `json:"baseProductId" `
	BaseProductUse int    `json:"baseProductUse" `
}

type UpdateProductI struct {
	Id            string                     `json:"id" validate:"required"`
	ProductName   string                     `json:"productName"`
	ProductNameTh string                     `json:"productNameTh"`
	BarcodeId     string                     `json:"barcodeId"`
	GroupId       string                     `json:"groupId"`
	BasePrice     float64                    `json:"basePrice"`
	SalePrice     float64                    `json:"salePrice"`
	Unit          string                     `json:"unit"`
	IsBaseProduct bool                       `json:"isBaseProduct"`
	SubProduct    []UpdateProductISubProduct `json:"subProduct"`
}

type UpdateProductISubProduct struct {
	Id             string `json:"id" validate:"required"`
	BaseProductId  string `json:"baseProductId"`
	BaseProductUse int    `json:"baseProductUse"`
}

type UpdateProductO struct{}

type DeleteProductI struct {
	Id string `json:"id" validate:"required"`
}

type DeleteProductO struct{}
