package tb

func (TProductGroup) TableName() string {
	return `t_product_group`
}

func (TProductGroup) Columns() []string {
	return []string{`id`, `group_name`}
}

type TProductGroup struct {
	Id        string `json:"id" gorm:"column:id;primaryKey" xorm:"id"`
	GroupName string `json:"group_name" gorm:"column:group_name;not null" xorm:"group_name"`

	TProduct []TProduct `gorm:"foreignKey:GroupId;references:Id" json:"-"`
}
