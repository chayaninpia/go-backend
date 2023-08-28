package collectionmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (Address) CollectionName() string {
	return `Address`
}

type BaseModel struct {
	NameTh *string `json:"name_th" bson:"name_th,omitempty"`
	NameEn *string `json:"name_en" bson:"name_en,omitempty"`
}

type BaseAmphureModel struct {
	NameTh   *string    `json:"name_th" bson:"name_th,omitempty"`
	NameEn   *string    `json:"name_en" bson:"name_en,omitempty"`
	Province *BaseModel `json:"province" bson:"province,omitempty"`
}

type Address struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	PostCode  *int               `json:"post_code" bson:"post_code,omitempty"`
	NameTh    *string            `json:"name_th" bson:"name_th,omitempty"`
	NameEn    *string            `json:"name_en" bson:"name_en,omitempty"`
	CreatedAt *time.Time         `json:"-" bson:"created_at,omitempty"`
	UpdatedAt *time.Time         `json:"-" bson:"updated_at,omitempty"`
	Amphure   *BaseAmphureModel  `json:"amphure" bson:"amphure,omitempty"`
}
