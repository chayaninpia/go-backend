package tablemodels

func (TProductStock) TableName() string {
	return `t_product_stock`
}

func (TProductStock) Columns() []string {
	return []string{`product_id`, `quantity`}
}

type TProductStock struct {
	ProductId string `json:"product_id" gorm:"column:product_id;primaryKey" xorm:"product_id"`
	Quantity  int    `json:"quantity" gorm:"column:quantity;not null" xorm:"quantity"`
}
