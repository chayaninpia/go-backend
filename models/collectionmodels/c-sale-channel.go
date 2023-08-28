package collectionmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (SaleChannel) CollectionName() string {
	return `Customer`
}

type SaleChannel struct {
	Id          primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	ShopId      *string            `json:"shopId" bson:"shop_id,omitempty"`
	SaleChannel *string            `json:"SaleChannel"  bson:"sale_channel,omitempty"`
	AppSecret   *string            `json:"appSecret" bson:"app_secret,omitempty"`
	AccessToken *string            `json:"accessToken" bson:"access_token,omitempty"`
	Sign        *string            `json:"sign" bson:"sign,omitempty"`

	CreatedAt *time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
