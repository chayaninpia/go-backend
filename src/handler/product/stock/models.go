package stock

type GetStockI struct {
	BarcodeId   string `json:"barcodeId"`
	ProductId   string `json:"productId"`
	ProductName string `json:"productName"`
}

type GetStockO struct {
	ProductId   string `json:"productId" `
	BarcodeId   string `json:"barcodeId"`
	ProductName string `json:"productName" `
	Quantity    int    `json:"quantity" `
}

type AddStockI struct {
	ProductId  string `json:"productId" `
	BarcodeId  string `json:"barcodeId"  validate:"required"`
	UpdateType string `json:"updateType"  validate:"required"`
	Quantity   int    `json:"quantity"  validate:"required"`
}

type AddStockO struct{}
