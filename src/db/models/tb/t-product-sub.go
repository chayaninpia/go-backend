package tb

func (TProductSub) TableName() string {
	return `t_product_sub`
}

func (TProductSub) Columns() []string {
	return []string{`id`, `product_id`, `base_product_id`, `base_product_use`}
}

type TProductSub struct {
	Id              string `json:"id" gorm:"column:id;primaryKey" xorm:"id"`
	ProductId       string `json:"product_id" gorm:"column:product_id;not null" xorm:"product_id"`
	BaseProductId   string `json:"base_product_id" gorm:"base_product_id;not null" xorm:"base_product_id"`
	BaseProductName string `json:"base_product_name" gorm:"base_product_name;not null" xorm:"base_product_name"`
	BaseProductUse  int    `json:"base_product_use" gorm:"column:base_product_use;not null" xorm:"base_product_use"`
}
