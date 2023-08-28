package tablemodels

func (TProduct) TableName() string {
	return `t_product`
}

func (TProduct) Columns() []string {
	return []string{`id`, `product_name`, `product_name_th`, `barcode_id`, `group_id`, `base_price`, `sale_price`, `unit`, `is_base_product`}
}

type TProduct struct {
	Id            string  `json:"id" gorm:"column:id;primaryKey" xorm:"id"`
	ProductName   string  `json:"product_name" gorm:"column:product_name;not null" xorm:"product_name"`
	ProductNameTh string  `json:"product_name_th" gorm:"product_name_th" xorm:"product_name_th"`
	BarcodeId     string  `json:"barcode_id" gorm:"column:barcode_id;not null" xorm:"barcode_id"`
	GroupId       string  `json:"group_id" gorm:"column:group_id" xorm:"group_id"`
	BasePrice     float64 `json:"base_price" gorm:"column:base_price;not null" xorm:"base_price"`
	SalePrice     float64 `json:"sale_price" gorm:"column:sale_price;not null" xorm:"sale_price"`
	Unit          string  `json:"unit" gorm:"column:unit;not null" xorm:"unit"`
	IsBaseProduct bool    `json:"is_base_product" gorm:"column:is_base_product;not null" xorm:"is_base_product"`

	TProductStock   []TProductStock `gorm:"foreignKey:ProductId;references:Id" json:"-"`
	TProductSubMain []TProductSub   `gorm:"foreignKey:ProductId;references:Id" json:"-"`
	TProductSubBase []TProductSub   `gorm:"foreignKey:BaseProductId;references:Id" json:"-"`
}
