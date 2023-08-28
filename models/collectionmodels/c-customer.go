package collectionmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (Customer) CollectionName() string {
	return `Customer`
}

type Customer struct {
	Id        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	FullName  *string            `json:"full_name"  bson:"full_name,omitempty"`
	ContactNo *string            `json:"contact_no"  bson:"contact_no,omitempty"`
	Email     *string            `json:"email"  bson:"email,omitempty"`
	Address   *string            `json:"address"  bson:"address,omitempty"`
	Tambon    *string            `json:"tambon" bson:"tambon,omitempty"`
	Amphure   *string            `json:"amphure"  bson:"amphure,omitempty"`
	Province  *string            `json:"province"  bson:"province,omitempty"`
	PostCode  *int               `json:"post_code"  bson:"post_code,omitempty"`
	CreatedAt *time.Time         `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt *time.Time         `json:"updated_at" bson:"updated_at,omitempty"`
}
