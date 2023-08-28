package models

import "apps/models/collectionmodels"

type SaleOrderCreateI struct {
	SaleType         string                        `json:"sale_type"  bson:"sale_type,omitempty"`                //ประเภทการขาย โอน ปลายทาง
	SaleChannel      string                        `json:"sale_channel"  bson:"sale_channel,omitempty"`          //ช่องทางการขาย shopee laz lineOA
	ShopId           string                        `json:"shop_id" bson:"shop_id,omitempty"`                     //รหัสร้านค้า
	ItemList         []collectionmodels.ItemList   `json:"item_list" bson:"item_list,omitempty"`                 //สินค้าที่ขาย
	OrderAddress     collectionmodels.OrderAddress `json:"order_address" bson:"order_address,omitempty"`         //ที่อยู่ผู้ซื้อ
	ShippingPlatform string                        `json:"shipping_platform" bson:"shipping_platform,omitempty"` //ผู้จัดส่ง kerry flash
	ShippingFee      float64                       `json:"shipping_fee" bson:"shipping_fee,omitempty"`           //ค่าขนส่ง
	TrackingCode     string                        `json:"tracking_code" bson:"tracking_code,omitempty"`         //เลข tracking
}

type SaleOrderCreateO struct {
	OrderNo          string                        `json:"order_no"  bson:"order_no,omitempty"`
	SaleType         string                        `json:"sale_type"  bson:"sale_type,omitempty"`                //ประเภทการขาย โอน ปลายทาง
	SaleChannel      string                        `json:"sale_channel"  bson:"sale_channel,omitempty"`          //ช่องทางการขาย shopee laz lineOA
	ShopId           string                        `json:"shop_id" bson:"shop_id,omitempty"`                     //รหัสร้านค้า
	TotalBasePrice   float64                       `json:"total_base_price"  bson:"total_base_price,omitempty"`  //ต้นทุน
	TotalSalePrice   float64                       `json:"total_sale_price"  bson:"total_sale_price,omitempty"`  //ราคาที่ขาย
	TotalVat         float64                       `json:"total_vat" bson:"total_vat,omitempty"`                 //จำนวนภาษี
	TotalProfit      float64                       `json:"total_profit"  bson:"total_profit,omitempty"`          //กำไร
	ItemList         []collectionmodels.ItemList   `json:"item_list" bson:"item_list,omitempty"`                 //สินค้าที่ขาย
	OrderAddress     collectionmodels.OrderAddress `json:"order_address" bson:"order_address,omitempty"`         //ที่อยู่ผู้ซื้อ
	ShippingPlatform string                        `json:"shipping_platform" bson:"shipping_platform,omitempty"` //ผู้จัดส่ง kerry flash
	ShippingFee      float64                       `json:"shipping_fee" bson:"shipping_fee,omitempty"`           //ค่าขนส่ง
	TrackingCode     string                        `json:"tracking_code" bson:"tracking_code,omitempty"`         //เลข tracking
	Status           string                        `json:"status" bson:"status,omitempty"`                       //สถานะ order
}

type SaleOrderReadI struct {
	OrderNo string `json:"orderNo"`
}

type SaleOrderReadO SaleOrderCreateO

type SaleOrderUpdateI struct {
	OrderNo      string `json:"order_no"  bson:"order_no,omitempty"`
	TrackingCode string `json:"tracking_code" bson:"tracking_code,omitempty"` //เลข tracking
	Status       string `json:"status" bson:"status,omitempty"`               //สถานะ order
}

type SaleOrderUpdateO SaleOrderCreateO
