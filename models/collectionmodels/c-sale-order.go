package collectionmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (SaleOrder) CollectionName() string {
	return `SaleOrder`
}

type SaleOrder struct {
	Id               primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	OrderNo          *string            `json:"order_no"  bson:"order_no,omitempty"`
	SaleType         *string            `json:"sale_type"  bson:"sale_type,omitempty"`                //ประเภทการขาย โอน ปลายทาง
	SaleChannel      *string            `json:"sale_channel"  bson:"sale_channel,omitempty"`          //ช่องทางการขาย shopee laz lineOA
	ShopId           *string            `json:"shop_id" bson:"shop_id,omitempty"`                     //รหัสร้านค้า
	TotalBasePrice   *float64           `json:"total_base_price"  bson:"total_base_price,omitempty"`  //ต้นทุน
	TotalSalePrice   *float64           `json:"total_sale_price"  bson:"total_sale_price,omitempty"`  //ราคาที่ขาย
	TotalVat         *float64           `json:"total_vat" bson:"total_vat,omitempty"`                 //จำนวนภาษี
	TotalProfit      *float64           `json:"total_profit"  bson:"total_profit,omitempty"`          //กำไร
	ItemList         *[]ItemList        `json:"item_list" bson:"item_list,omitempty"`                 //สินค้าที่ขาย
	OrderAddress     *OrderAddress      `json:"order_address" bson:"order_address,omitempty"`         //ที่อยู่ผู้ซื้อ
	ShippingPlatform *string            `json:"shipping_platform" bson:"shipping_platform,omitempty"` //ผู้จัดส่ง kerry flash
	ShippingFee      *float64           `json:"shipping_fee" bson:"shipping_fee,omitempty"`           //ค่าขนส่ง
	TrackingCode     *string            `json:"tracking_code" bson:"tracking_code,omitempty"`         //เลข tracking
	Status           *string            `json:"status" bson:"status,omitempty"`                       //สถานะ order
	CreatedAt        *time.Time         `json:"order_date"  bson:"order_date,omitempty"`
	UpdatedAt        *time.Time         `json:"updated_at" bson:"updated_at,omitempty"`
}

type OrderAddress struct {
	Country   *string `json:"country" bson:"country,omitempty"`
	Address3  *string `json:"address3" bson:"address3,omitempty"`
	Address2  *string `json:"address2" bson:"address2,omitempty"`
	City      *string `json:"city" bson:"city,omitempty"`
	Phone     *string `json:"phone" bson:"phone,omitempty"`
	Address1  *string `json:"address1" bson:"address1,omitempty"`
	PostCode  *string `json:"post_code" bson:"post_code,omitempty"`
	Phone2    *string `json:"phone2" bson:"phone2,omitempty"`
	LastName  *string `json:"last_name" bson:"last_name,omitempty"`
	Address5  *string `json:"address5" bson:"address5,omitempty"`
	Address4  *string `json:"address4" bson:"address4,omitempty"`
	FirstName *string `json:"first_name" bson:"first_name,omitempty"`
}

type ItemList struct {
	ProductId *string  `json:"productId" bson:"product_id,omitempty"`
	ItemId    *string  `json:"itemId" bson:"item_id,omitempty"`
	SellerSku *string  `json:"sellerSku" bson:"seller_sku,omitempty"`
	Quantity  *int     `json:"quantity" bson:"quantity,omitempty"`
	BasePrice *float64 `json:"base_price"  bson:"base_price,omitempty"`
	SalePrice *float64 `json:"sale_price"  bson:"sale_price,omitempty"`
	Vat       *float64 `json:"vat" bson:"vat,omitempty"`
}
