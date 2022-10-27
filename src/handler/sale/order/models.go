package order

import "time"

type CreateSaleOrderI struct {
	SaleType      string                 `json:"saleType" validate:"required"`
	SaleChannelId string                 `json:"saleChannelId" validate:"required"`
	OtherCost     float64                `json:"otherCost" `
	Items         []CreateSaleOrderIItem `json:"items" validate:"required"`
}

type CreateSaleOrderIItem struct {
	ProductId string  `json:"productId" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	BasePrice float64 `json:"basePrice" validate:"required"`
	SalePrice float64 `json:"salePrice" validate:"required"`
}

type CreateSaleOrderO struct {
	OrderNo        string    `json:"orderNo" `
	SaleType       string    `json:"saleType" `
	SaleChannelId  string    `json:"saleChannelId" `
	TotalBasePrice float64   `json:"totalBasePrice" `
	TotalSalePrice float64   `json:"totalSalePrice" `
	OtherCost      float64   `json:"otherCost" `
	TotalProfit    float64   `json:"totalProfit" `
	OrderDate      time.Time `json:"orderDate" `
}

type ReadSaleOrderI struct {
	OrderNo string `json:"orderNo"`
}

type ReadSaleOrderO struct {
	OrderNo        string               `json:"orderNo" `
	SaleType       string               `json:"saleType" `
	SaleChannelId  string               `json:"saleChannelId" `
	TotalBasePrice float64              `json:"totalBasePrice" `
	TotalSalePrice float64              `json:"totalSalePrice" `
	OtherCost      float64              `json:"otherCost" `
	TotalProfit    float64              `json:"totalProfit" `
	OrderDate      time.Time            `json:"orderDate" `
	Items          []ReadSaleOrderOItem `json:"items"`
}

type ReadSaleOrderOItem struct {
	ProductId string  `json:"productId"`
	Quantity  int     `json:"quantity"`
	BasePrice float64 `json:"basePrice"`
	SalePrice float64 `json:"salePrice"`
}
