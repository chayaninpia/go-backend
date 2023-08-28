package tablemodels

import "time"

func (TDSubDistricts) TableName() string {
	return `t_sub_districts`
}

func (TDSubDistricts) Columns() []string {
	return []string{"id", "name_th", "name_en", "district_id", "post_code", "created_at", "updated_at"}
}

type TDSubDistricts struct {
	Id         string    `json:"id" gorm:"id;primaryKey" xorm:"id"`
	NameTh     string    `json:"name_th" gorm:"name_th;not null" xorm:"name_th"`
	NameEn     string    `json:"name_en" gorm:"name_en;not null" xorm:"name_en"`
	DistrictId string    `json:"district_id" gorm:"district_id;not null" xorm:"district_id"`
	PostCode   string    `json:"post_code" gorm:"post_code;not null" xorm:"post_code"`
	CreatedAt  time.Time `json:"created_at" gorm:"created_at;not null" xorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at;not null" xorm:"updated_at"`
}
