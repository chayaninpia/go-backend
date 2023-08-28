package tablemodels

func (TSaleChannel) TableName() string {
	return `t_sale_channel`
}

func (TSaleChannel) Columns() []string {
	return []string{`id`, `sale_channel`}
}

type TSaleChannel struct {
	Id          string `json:"id" gorm:"id;primaryKey" xorm:"id"`
	SaleChannel string `json:"sale_channel" gorm:"sale_channel;not null" xorm:"sale_channel"`

	TSaleOrder []TSaleOrder `gorm:"foreignKey:SaleChannelId;references:Id" json:"-"`
}
