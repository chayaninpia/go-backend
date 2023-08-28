package tablemodels

func (TCustomer) TableName() string {
	return `t_customer`
}

func (TCustomer) Columns() []string {
	return []string{`id`, `full_name`, `contact_no`, `email`, `address`, `sub_district`, `province`, `post_code`}
}

type TCustomer struct {
	Id          string `json:"id" gorm:"id;primaryKey" xorm:"id"`
	FullName    string `json:"full_name" gorm:"full_name;not null" xorm:"full_name"`
	ContactNo   string `json:"contact_no" gorm:"contact_no;not null" xorm:"contact_no"`
	Email       string `json:"email" gorm:"email;not null" xorm:"email"`
	Address     string `json:"address" gorm:"address;not null" xorm:"address"`
	SubDistrict string `json:"sub_district" gorm:"sub_district" xorm:"sub_district"`
	District    string `json:"district" gorm:"district;not null" xorm:"district"`
	Province    string `json:"province" gorm:"province;not null" xorm:"province"`
	PostCode    string `json:"post_code" gorm:"post_code;not null" xorm:"post_code"`
}
