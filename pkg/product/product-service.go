package product

import (
	"apps/models"
	"apps/models/collectionmodels"
	"apps/pkg/utils/timex"
	"apps/pkg/utils/tox"
	"apps/pkg/utils/validx"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProductService interface {
	Create(req models.ProductCreateI) error
	Read(req models.ProductReadI) ([]models.ProductReadO, error)
	ReadAll() ([]models.ProductReadO, error)
	ReadById(req models.ProductReadI) (*models.ProductReadO, error)
	Update(req models.ProductUpdateI) (*models.UpdateProductO, error)
	Delete(req models.ProductDeleteI) error
}

type ProductService struct {
	ProductRepository IProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{
		ProductRepository: NewProductRepository(),
	}
}

func (service ProductService) Create(req models.ProductCreateI) error {

	if err := validx.Struct(req); err != nil {
		return err
	}

	timeNow := timex.TimeNowPtr()
	inputData := collectionmodels.Product{
		Id:          primitive.NewObjectID(),
		ProductId:   &req.ProductId,
		NameTh:      &req.NameTh,
		NameEn:      &req.NameEn,
		Images:      &req.Images,
		SkuId:       &req.SkuId,
		ShopItem:    &req.ShopItem,
		ProductList: &req.ProductList,
		ShopSku:     &req.ShopSku,
		URL:         &req.URL,
		Stock:       &req.Stock,
		Price:       &req.Price,
		Status:      &req.Status,
		CreatedAt:   timeNow,
		UpdatedAt:   timeNow,
	}

	if err := service.ProductRepository.Create(inputData); err != nil {
		return err
	}

	return nil
}

func (service ProductService) Read(req models.ProductReadI) ([]models.ProductReadO, error) {

	inputData := collectionmodels.Product{
		NameTh: &req.NameTh,
		NameEn: &req.NameEn,
	}

	resMain, err := service.ProductRepository.Read(inputData)
	if err != nil {
		return nil, err
	}

	res := make([]models.ProductReadO, 0)
	for _, v := range resMain {
		res = append(res, models.ProductReadO{
			Id:          v.Id.Hex(),
			ProductId:   tox.String(v.ProductId),
			NameTh:      tox.String(v.NameTh),
			NameEn:      tox.String(v.NameEn),
			Images:      tox.Slice[string](v.Images),
			SkuId:       tox.String(v.SkuId),
			ShopItem:    tox.Struct[collectionmodels.ShopItem](v.ShopItem),
			ProductList: tox.Slice[collectionmodels.ProductList](v.ProductList),
			ShopSku:     tox.String(v.ShopSku),
			URL:         tox.String(v.URL),
			Stock:       tox.Struct[collectionmodels.Stock](v.Stock),
			Price:       tox.Struct[collectionmodels.Price](v.Price),
			Status:      tox.Bool(v.Status),
		})
	}

	return res, nil
}

func (service ProductService) ReadAll() ([]models.ProductReadO, error) {

	resMain, err := service.ProductRepository.ReadAll()
	if err != nil {
		return nil, err
	}

	res := make([]models.ProductReadO, 0)
	for _, v := range resMain {
		res = append(res, models.ProductReadO{
			Id:          v.Id.Hex(),
			ProductId:   tox.String(v.ProductId),
			NameTh:      tox.String(v.NameTh),
			NameEn:      tox.String(v.NameEn),
			Images:      tox.Slice[string](v.Images),
			SkuId:       tox.String(v.SkuId),
			ShopItem:    tox.Struct[collectionmodels.ShopItem](v.ShopItem),
			ProductList: tox.Slice[collectionmodels.ProductList](v.ProductList),
			ShopSku:     tox.String(v.ShopSku),
			URL:         tox.String(v.URL),
			Stock:       tox.Struct[collectionmodels.Stock](v.Stock),
			Price:       tox.Struct[collectionmodels.Price](v.Price),
			Status:      tox.Bool(v.Status),
		})
	}

	return res, nil
}

func (service ProductService) ReadById(req models.ProductReadI) (*models.ProductReadO, error) {

	id, _ := primitive.ObjectIDFromHex(req.Id)
	inputData := collectionmodels.Product{
		Id: id,
	}

	resMain, err := service.ProductRepository.ReadById(inputData)
	if err != nil {
		return nil, err
	}

	res := models.ProductReadO{
		Id:          resMain.Id.Hex(),
		ProductId:   tox.String(resMain.ProductId),
		NameTh:      tox.String(resMain.NameTh),
		NameEn:      tox.String(resMain.NameEn),
		Images:      tox.Slice[string](resMain.Images),
		SkuId:       tox.String(resMain.SkuId),
		ShopItem:    tox.Struct[collectionmodels.ShopItem](resMain.ShopItem),
		ProductList: tox.Slice[collectionmodels.ProductList](resMain.ProductList),
		ShopSku:     tox.String(resMain.ShopSku),
		URL:         tox.String(resMain.URL),
		Stock:       tox.Struct[collectionmodels.Stock](resMain.Stock),
		Price:       tox.Struct[collectionmodels.Price](resMain.Price),
		Status:      tox.Bool(resMain.Status),
	}

	return &res, nil
}

func (service ProductService) Update(req models.ProductUpdateI) (*models.UpdateProductO, error) {

	if err := validx.Struct(req); err != nil {
		return nil, err
	}

	id, _ := primitive.ObjectIDFromHex(req.ProductId)
	updateData := collectionmodels.Product{
		Id:          id,
		ProductId:   &req.ProductId,
		NameTh:      &req.NameTh,
		NameEn:      &req.NameEn,
		Images:      &req.Images,
		SkuId:       &req.SkuId,
		ShopItem:    &req.ShopItem,
		ProductList: &req.ProductList,
		ShopSku:     &req.ShopSku,
		URL:         &req.URL,
		Stock:       &req.Stock,
		Price:       &req.Price,
		Status:      &req.Status,
		UpdatedAt:   timex.TimeNowPtr(),
	}

	res, err := service.ProductRepository.Update(updateData)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service ProductService) Delete(req models.ProductDeleteI) error {

	if err := validx.Struct(req); err != nil {
		return err
	}

	id, _ := primitive.ObjectIDFromHex(req.Id)
	deleteFilter := collectionmodels.Product{Id: id}

	if err := service.ProductRepository.Delete(deleteFilter); err != nil {
		return err
	}

	return nil
}

func (service ProductService) AddStock(req models.ProductUpdateStockI) (*models.ProductUpdateStockO, error) {

	if err := validx.Struct(req); err != nil {
		return nil, err
	}

	getProduct := models.ProductReadI{
		Id: req.Id,
	}

	productInfo, err := service.ReadById(getProduct)
	if err != nil {
		return nil, err
	}

	totalQuantity := tox.Int(productInfo.Stock.Quantity) + *req.Stock.Quantity

	id, _ := primitive.ObjectIDFromHex(req.Id)
	updateInfo := collectionmodels.Product{
		Id: id,
		Stock: &collectionmodels.Stock{
			Quantity: &totalQuantity,
		},
	}

	if _, err := service.ProductRepository.UpdateStock(updateInfo); err != nil {
		return nil, err
	}

	return &models.ProductUpdateStockO{
		Id:    req.Id,
		Stock: tox.Struct[collectionmodels.Stock](updateInfo.Stock),
	}, nil
}
