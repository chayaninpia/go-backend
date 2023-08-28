package models

import "apps/models/collectionmodels"

type ProductCreateI struct {
	ProductId   string                         `json:"productId" bson:"product_id"`
	NameTh      string                         `json:"nameTh" bson:"name_th"`
	NameEn      string                         `json:"nameEn" bson:"name_en"`
	Images      []string                       `json:"images" bson:"images"`
	SkuId       string                         `json:"skuId" bson:"sku_id"`
	ShopItem    collectionmodels.ShopItem      `json:"shopItem" bson:"shop_item"`
	ProductList []collectionmodels.ProductList `json:"productList" bson:"product_list"`
	ShopSku     string                         `json:"shopSku" bson:"shop_sku"`
	URL         string                         `json:"url" bson:"url"`
	Stock       collectionmodels.Stock         `json:"stock" bson:"stock"`
	Price       collectionmodels.Price         `json:"price" bson:"price"`
	Status      bool                           `json:"status" bson:"status"`
}

type ProductCreateO struct{}

type ProductReadI struct {
	Id     string `json:"id" `
	NameTh string `json:"nameTh" bson:"name_th"`
	NameEn string `json:"nameEn" bson:"name_en"`
}

type ProductReadO struct {
	Id          string                         `json:"id" bson:"id"`
	ProductId   string                         `json:"productId" bson:"product_id"`
	NameTh      string                         `json:"nameTh" bson:"name_th"`
	NameEn      string                         `json:"nameEn" bson:"name_en"`
	Images      []string                       `json:"images" bson:"images"`
	SkuId       string                         `json:"skuId" bson:"sku_id"`
	ShopItem    collectionmodels.ShopItem      `json:"shopItem" bson:"shop_item"`
	ProductList []collectionmodels.ProductList `json:"productList" bson:"product_list"`
	ShopSku     string                         `json:"shopSku" bson:"shop_sku"`
	URL         string                         `json:"url" bson:"url"`
	Stock       collectionmodels.Stock         `json:"stock" bson:"stock"`
	Price       collectionmodels.Price         `json:"price" bson:"price"`
	Status      bool                           `json:"status" bson:"status"`
}

type ProductUpdateI struct {
	Id          string                         `json:"id" bson:"id"`
	ProductId   string                         `json:"productId" bson:"product_id"`
	NameTh      string                         `json:"nameTh" bson:"name_th"`
	NameEn      string                         `json:"nameEn" bson:"name_en"`
	Images      []string                       `json:"images" bson:"images"`
	SkuId       string                         `json:"skuId" bson:"sku_id"`
	ShopItem    collectionmodels.ShopItem      `json:"shopItem" bson:"shop_item"`
	ProductList []collectionmodels.ProductList `json:"productList" bson:"product_list"`
	ShopSku     string                         `json:"shopSku" bson:"shop_sku"`
	URL         string                         `json:"url" bson:"url"`
	Stock       collectionmodels.Stock         `json:"stock" bson:"stock"`
	Price       collectionmodels.Price         `json:"price" bson:"price"`
	Status      bool                           `json:"status" bson:"status"`
}

type ProductUpdateO ProductUpdateI

type ProductDeleteI struct {
	Id string `json:"id" validate:"required"`
}

type ProductDeleteO struct{}

type ProductUpdateStockI struct {
	Id    string                 `json:"id" validate:"required"`
	Stock collectionmodels.Stock `json:"stock" bson:"stock"`
}

type ProductUpdateStockO struct {
	Id    string                 `json:"id" validate:"required"`
	Stock collectionmodels.Stock `json:"stock" bson:"stock"`
}
