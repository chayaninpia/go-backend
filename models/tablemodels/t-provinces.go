package tablemodels

import "time"

func (TProvinces) TableName() string {
	return `t_provinces`
}

func (TProvinces) Columns() []string {
	return []string{"id", "name_th", "name_en", "geography_id", "created_at", "updated_at"}
}

type TProvinces struct {
	Id          string    `json:"id" gorm:"id;primaryKey" xorm:"id"`
	NameTh      string    `json:"name_th" gorm:"name_th;not null" xorm:"name_th"`
	NameEn      string    `json:"name_en" gorm:"name_en;not null" xorm:"name_en"`
	GeographyId string    `json:"geography_id" gorm:"geography_id;not null" xorm:"geography_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at;not null" xorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at;not null" xorm:"updated_at"`

	TDistricts []TDistricts `gorm:"foreignKey:ProvinceId;references:Id" json:"-"`
}
