package tablemodels

func (TGeographies) TableName() string {
	return `t_geographies`
}

func (TGeographies) Columns() []string {
	return []string{`id`, `name`}
}

type TGeographies struct {
	Id   string `json:"id" gorm:"id;primaryKey" xorm:"id"`
	Name string `json:"name" gorm:"name;not null" xorm:"name"`

	TProvinces []TProvinces `gorm:"foreignKey:GeographyId;references:Id" json:"-"`
}
