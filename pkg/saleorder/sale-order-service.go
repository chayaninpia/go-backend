package saleorder

import (
	"apps/models"
	"apps/models/collectionmodels"
	"apps/pkg/constants"
	"apps/pkg/product"
	"apps/pkg/utils/timex"
	"apps/pkg/utils/tox"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ISaleOrderService interface {
	Create(req models.SaleOrderCreateI) error
	Read() ([]models.SaleOrderReadO, error)
	ReadByOrderNo(req models.SaleOrderReadI) (*models.SaleOrderReadO, error)
	Update(req models.SaleOrderUpdateI) error
}

type SaleOrderService struct {
	SaleOrderRepository ISaleOrderRepository
	ProductRepository   product.IProductRepository
}

func NewSaleOrderService() *SaleOrderService {
	return &SaleOrderService{
		SaleOrderRepository: NewSaleOrderRepository(),
		ProductRepository:   product.NewProductRepository(),
	}
}

func (service *SaleOrderService) Create(req models.SaleOrderCreateI) error {

	timeNow := timex.TimeNowPtr()
	orderNo := strconv.Itoa(int(time.Now().Unix()))
	productIds := make([]string, 0)
	for _, v := range req.ItemList {
		productIds = append(productIds, tox.String(v.ProductId))
	}
	listProduct, err := service.ProductRepository.ReadInProductId(productIds)
	if err != nil {
		return err
	}

	var totalBasePrice, TotalSalePrice, TotalVat, TotalProfit float64
	for _, v := range listProduct {
		totalBasePrice += tox.Float(v.Price.BasePrice)
		TotalSalePrice += tox.Float(v.Price.Price)
		TotalVat += tox.Float(v.Price.Vat)
		TotalProfit = TotalSalePrice - totalBasePrice - TotalVat
	}

	toInsert := collectionmodels.SaleOrder{
		Id:               primitive.NewObjectID(),
		OrderNo:          &orderNo,
		SaleType:         &req.SaleType,
		SaleChannel:      &req.SaleChannel,
		ShopId:           &req.ShopId,
		TotalBasePrice:   &totalBasePrice,
		TotalSalePrice:   &TotalSalePrice,
		TotalVat:         &TotalVat,
		TotalProfit:      &TotalProfit,
		ItemList:         &req.ItemList,
		OrderAddress:     &req.OrderAddress,
		ShippingPlatform: &req.ShippingPlatform,
		ShippingFee:      &req.ShippingFee,
		TrackingCode:     &req.TrackingCode,
		Status:           &constants.OrderStatus.Unpaid,
		CreatedAt:        timeNow,
		UpdatedAt:        timeNow,
	}
	if err := service.SaleOrderRepository.Create(toInsert); err != nil {
		return err
	}
	return nil
}

func (service *SaleOrderService) Read() ([]models.SaleOrderReadO, error) {

	result, err := service.SaleOrderRepository.Read()
	if err != nil {
		return nil, err
	}

	res := make([]models.SaleOrderReadO, 0)
	for _, v := range result {
		res = append(res, models.SaleOrderReadO{
			OrderNo:          tox.String(v.OrderNo),
			SaleType:         tox.String(v.SaleType),
			SaleChannel:      tox.String(v.SaleChannel),
			ShopId:           tox.String(v.ShopId),
			TotalBasePrice:   tox.Float(v.TotalBasePrice),
			TotalSalePrice:   tox.Float(v.TotalSalePrice),
			TotalVat:         tox.Float(v.TotalVat),
			TotalProfit:      tox.Float(v.TotalProfit),
			ItemList:         tox.Slice[collectionmodels.ItemList](v.ItemList),
			OrderAddress:     tox.Struct[collectionmodels.OrderAddress](v.OrderAddress),
			ShippingPlatform: tox.String(v.ShippingPlatform),
			ShippingFee:      tox.Float(v.ShippingFee),
			TrackingCode:     tox.String(v.TrackingCode),
			Status:           tox.String(v.Status),
		})
	}

	return res, nil
}

func (service *SaleOrderService) ReadByOrderNo(req models.SaleOrderReadI) (*models.SaleOrderReadO, error) {

	toFilter := collectionmodels.SaleOrder{OrderNo: &req.OrderNo}
	result, err := service.SaleOrderRepository.ReadByOrderNo(toFilter)
	if err != nil {
		return nil, err
	}
	res := &models.SaleOrderReadO{
		OrderNo:          tox.String(result.OrderNo),
		SaleType:         tox.String(result.SaleType),
		SaleChannel:      tox.String(result.SaleChannel),
		ShopId:           tox.String(result.ShopId),
		TotalBasePrice:   tox.Float(result.TotalBasePrice),
		TotalSalePrice:   tox.Float(result.TotalSalePrice),
		TotalVat:         tox.Float(result.TotalVat),
		TotalProfit:      tox.Float(result.TotalProfit),
		ItemList:         tox.Slice[collectionmodels.ItemList](result.ItemList),
		OrderAddress:     tox.Struct[collectionmodels.OrderAddress](result.OrderAddress),
		ShippingPlatform: tox.String(result.ShippingPlatform),
		ShippingFee:      tox.Float(result.ShippingFee),
		TrackingCode:     tox.String(result.TrackingCode),
		Status:           tox.String(result.Status),
	}

	return res, nil
}

func (service *SaleOrderService) Update(req models.SaleOrderUpdateI) error {

	toUpdate := collectionmodels.SaleOrder{OrderNo: &req.OrderNo, UpdatedAt: timex.TimeNowPtr()}
	if len(req.TrackingCode) > 0 {
		toUpdate.TrackingCode = &req.TrackingCode
	}
	if len(req.Status) > 0 {
		toUpdate.Status = &req.Status
	}

	if err := service.SaleOrderRepository.Update(toUpdate); err != nil {
		return err
	}

	return nil
}
