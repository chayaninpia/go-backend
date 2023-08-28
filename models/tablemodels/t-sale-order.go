package tablemodels

import "time"

func (TSaleOrder) TableName() string {
	return `t_sale_order`
}

func (TSaleOrder) Columns() []string {
	return []string{`order_no`, `sale_type`, `sale_channel_id`, `total_base_price`, `total_sale_price`, `other_cost`, `total_profit`, `order_date`}
}

type TSaleOrder struct {
	OrderNo        string    `json:"order_no" gorm:"order_no;primaryKey" xorm:"order_no"`
	CustomerId     string    `json:"customer_id" gorm:"customer_id" xorm:"customer_id"`
	SaleType       string    `json:"sale_type" gorm:"sale_type;not null" xorm:"sale_type"`
	SaleChannelId  string    `json:"sale_channel_id" gorm:"sale_channel_id;not null" xorm:"sale_channel_id"`
	TotalBasePrice float64   `json:"total_base_price" gorm:"total_base_price;not null" xorm:"total_base_price"`
	TotalSalePrice float64   `json:"total_sale_price" gorm:"total_sale_price;not null" xorm:"total_sale_price"`
	OtherCost      float64   `json:"other_cost" gorm:"other_cost" xorm:"other_cost"`
	TotalProfit    float64   `json:"total_profit" gorm:"total_profit;not null" xorm:"total_profit"`
	OrderDate      time.Time `json:"order_date" gorm:"order_date;not null" xorm:"order_date"`

	TSaleOrderItem []TSaleOrderItem `gorm:"foreignKey:OrderNo;references:OrderNo" json:"-"`
}
