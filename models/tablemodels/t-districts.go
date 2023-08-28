package tablemodels

import "time"

func (TDistricts) TableName() string {
	return `t_districts`
}

func (TDistricts) Columns() []string {
	return []string{"id", "name_th", "name_en", "province_id", "created_at", "updated_at"}
}

type TDistricts struct {
	Id         string    `json:"id" gorm:"id;primaryKey" xorm:"id"`
	NameTh     string    `json:"name_th" gorm:"name_th;not null" xorm:"name_th"`
	NameEn     string    `json:"name_en" gorm:"name_en;not null" xorm:"name_en"`
	ProvinceId string    `json:"province_id" gorm:"province_id;not null" xorm:"province_id"`
	CreatedAt  time.Time `json:"created_at" gorm:"created_at;not null" xorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at;not null" xorm:"updated_at"`

	TDSubDistricts []TDSubDistricts `gorm:"foreignKey:DistrictId;references:Id" json:"-"`
}
