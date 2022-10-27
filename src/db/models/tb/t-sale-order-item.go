package tb

func (TSaleOrderItem) TableName() string {
	return `t_sale_order_item`
}

func (TSaleOrderItem) Columns() []string {
	return []string{`order_no`, `product_id`, `quantity`, `base_price`, `sale_price`}
}

type TSaleOrderItem struct {
	OrderNo   string  `json:"order_no" gorm:"order_no;primaryKey" xorm:"order_no"`
	ProductId string  `json:"product_id" gorm:"product_id;primaryKey;not null" xorm:"product_id"`
	Quantity  int     `json:"quantity" gorm:"quantity;not null" xorm:"quantity"`
	BasePrice float64 `json:"base_price" gorm:"base_price;not null" xorm:"base_price"`
	SalePrice float64 `json:"sale_price" gorm:"sale_price;not null" xorm:"sale_price"`
}
