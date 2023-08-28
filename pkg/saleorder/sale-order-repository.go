package saleorder

import (
	"apps/models/collectionmodels"
	"apps/pkg/utils/mongox"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ISaleOrderRepository interface {
	Create(req collectionmodels.SaleOrder) error
	Read() ([]collectionmodels.SaleOrder, error)
	ReadByOrderNo(req collectionmodels.SaleOrder) (*collectionmodels.SaleOrder, error)
	Update(req collectionmodels.SaleOrder) error
}

type SaleOrderRepository struct {
	DB *mongox.DBService
}

func NewSaleOrderRepository() *SaleOrderRepository {
	return &SaleOrderRepository{}
}

func (repo SaleOrderRepository) getCollection() *mongo.Collection {
	return repo.DB.Database.Collection(collectionmodels.SaleOrder{}.CollectionName())
}

func (repo *SaleOrderRepository) getMongoService() error {

	if repo.DB == nil {

		dbService := &mongox.DBService{
			DBName: viper.GetString("mongodb.dbname"),
		}

		if cache := dbService.GetCache(); cache != nil {
			repo.DB = cache
			return nil
		}

		if err := dbService.NewService(); err != nil {
			return err
		}
		repo.DB = dbService.GetCache()
	}
	return nil
}

func (repo *SaleOrderRepository) Create(req collectionmodels.SaleOrder) error {

	if err := repo.getMongoService(); err != nil {
		return err
	}

	if _, err := repo.DB.InsertOne(repo.getCollection(), req); err != nil {
		return err
	}

	return nil
}

func (repo *SaleOrderRepository) Read() ([]collectionmodels.SaleOrder, error) {

	if err := repo.getMongoService(); err != nil {
		return nil, err
	}

	res := make([]collectionmodels.SaleOrder, 0)
	if err := repo.DB.FindDatas(repo.getCollection(), res, nil); err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *SaleOrderRepository) ReadByOrderNo(req collectionmodels.SaleOrder) (*collectionmodels.SaleOrder, error) {

	if err := repo.getMongoService(); err != nil {
		return nil, err
	}

	filter := bson.M{"order_no": req.OrderNo}

	res := &collectionmodels.SaleOrder{}

	if err := repo.DB.FindFirstData(repo.getCollection(), res, filter); err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *SaleOrderRepository) Update(req collectionmodels.SaleOrder) error {

	if err := repo.getMongoService(); err != nil {
		return err
	}

	filter := bson.M{"order_no": req.OrderNo}

	if _, err := repo.DB.UpdateOneByScript(repo.getCollection(), filter, req); err != nil {
		return err
	}

	return nil
}
