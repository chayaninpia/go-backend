package collectionmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (Product) CollectionName() string {
	return `Product`
}

type Product struct {
	Id          primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	ProductId   *string            `json:"productId" bson:"product_id,omitempty"`
	NameTh      *string            `json:"nameTh" bson:"name_th,omitempty"`
	NameEn      *string            `json:"nameEn" bson:"name_en,omitempty"`
	Images      *[]string          `json:"images" bson:"images,omitempty"`
	SkuId       *string            `json:"skuId" bson:"sku_id,omitempty"`
	ShopItem    *ShopItem          `json:"shopItem" bson:"shop_item,omitempty"`
	ProductList *[]ProductList     `json:"productList" bson:"product_list,omitempty"`
	ShopSku     *string            `json:"shopSku" bson:"shop_sku,omitempty"`
	URL         *string            `json:"url" bson:"url,omitempty"`
	Stock       *Stock             `json:"stock" bson:"stock,omitempty"`
	Price       *Price             `json:"price" bson:"price,omitempty"`
	Status      *bool              `json:"status" bson:"status,omitempty"`
	CreatedAt   *time.Time         `json:"createdAt" bson:"created_at,omitempty"`
	UpdatedAt   *time.Time         `json:"updatedAt" bson:"updated_at,omitempty"`
}

type Stock struct {
	Quantity  *int `json:"quantity" bson:"quantity,omitempty"`
	Available *int `json:"available" bson:"available,omitempty"`
	Remaining *int `json:"remaining" bson:"remaining,omitempty"`
}

type Price struct {
	Price        *float64 `json:"price" bson:"price,omitempty"`
	SpecialPrice *float64 `json:"specialPrice" bson:"special_price,omitempty"`
	BasePrice    *float64 `json:"basePrice" bson:"base_price,omitempty"`
	Vat          *float64 `json:"vat" bson:"vat,omitempty"`
}

type ShopItem struct {
	Lazada *StandardShopItem `json:"lazada" bson:"lazada,omitempty"`
	Shopee *StandardShopItem `json:"shopee" bson:"shopee,omitempty"`
}

type StandardShopItem struct {
	ItemId    *string `json:"itemId" bson:"item_id,omitempty"`
	SellerSku *string `json:"sellerSku" bson:"seller_sku,omitempty"`
}

type ProductList struct {
	ProductId   *string `json:"productId" bson:"product_id,omitempty"`
	QuantityUse *int    `json:"quantityUse" bson:"quantity_use,omitempty"`
}
