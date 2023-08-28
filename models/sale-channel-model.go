package models

type CreateSaleChannelI struct {
	SaleChannel string `json:"saleChannel" validate:"required"`
}

type CreateSaleChannelO struct{}

type ReadSaleChannelI struct {
	SaleChannel string `json:"saleChannel"`
}

type ReadSaleChannelO struct {
	Id          string `json:"id"`
	SaleChannel string `json:"saleChannel"`
}

type UpdateSaleChannelI struct {
	Id          string `json:"id" validate:"required"`
	SaleChannel string `json:"saleChannel"`
}

type UpdateSaleChannelO struct{}

type DeleteSaleChannelI struct {
	Id string `json:"id" validate:"required"`
}

type DeleteSaleChannelO struct{}
