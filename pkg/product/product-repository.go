package product

import (
	"apps/models"
	"apps/models/collectionmodels"
	"apps/pkg/utils/mongox"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"xorm.io/xorm"
)

type IProductRepository interface {
	Create(cp collectionmodels.Product) error
	CreateMany(cp []collectionmodels.Product) error
	Read(cp collectionmodels.Product) ([]collectionmodels.Product, error)
	ReadAll() ([]collectionmodels.Product, error)
	ReadById(cp collectionmodels.Product) (*collectionmodels.Product, error)
	ReadByProductId(cp collectionmodels.Product) (*collectionmodels.Product, error)
	ReadInProductId(productIds []string) ([]collectionmodels.Product, error)
	Update(cp collectionmodels.Product) (*models.UpdateProductO, error)
	Delete(cp collectionmodels.Product) error
	UpdateStock(cp collectionmodels.Product) (*models.ProductUpdateStockO, error)
}

type ProductRepository struct {
	E  *xorm.Engine
	DB *mongox.DBService
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (repo *ProductRepository) getCollection() *mongo.Collection {
	return repo.DB.Database.Collection(collectionmodels.Product{}.CollectionName())
}

func (repo *ProductRepository) getMongoService() error {

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

func (repo ProductRepository) Create(cp collectionmodels.Product) error {

	if err := repo.getMongoService(); err != nil {
		return err
	}

	if _, err := repo.DB.InsertOne(repo.getCollection(), cp); err != nil {
		return err
	}

	return nil
}

func (repo ProductRepository) CreateMany(cp []collectionmodels.Product) error {

	if err := repo.getMongoService(); err != nil {
		return err
	}

	toInsert := make([]interface{}, 0)
	for _, v := range cp {
		toInsert = append(toInsert, v)
	}

	if _, err := repo.DB.InsertMany(repo.getCollection(), toInsert); err != nil {
		return err
	}

	return nil
}

func (repo ProductRepository) Read(cp collectionmodels.Product) ([]collectionmodels.Product, error) {

	if err := repo.getMongoService(); err != nil {
		return nil, err
	}

	filter := bson.M{
		"name_th": bson.M{"$regex": cp.NameTh},
		"name_en": bson.M{"$regex": cp.NameEn},
	}

	res := make([]collectionmodels.Product, 0)

	if err := repo.DB.FindDatas(repo.getCollection(), &res, filter); err != nil {
		return nil, err
	}

	return res, nil
}

func (repo ProductRepository) ReadAll() ([]collectionmodels.Product, error) {

	if err := repo.getMongoService(); err != nil {
		return nil, err
	}

	res := make([]collectionmodels.Product, 0)

	filter := bson.M{}

	if err := repo.DB.FindDatas(repo.getCollection(), &res, filter); err != nil {
		return nil, err
	}

	return res, nil
}

func (repo ProductRepository) ReadById(cp collectionmodels.Product) (*collectionmodels.Product, error) {

	if err := repo.getMongoService(); err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": cp.Id,
	}

	res := collectionmodels.Product{}

	if err := repo.DB.FindFirstData(repo.getCollection(), &res, filter); err != nil {
		return nil, err
	}

	return &res, nil
}

func (repo ProductRepository) ReadByProductId(cp collectionmodels.Product) (*collectionmodels.Product, error) {

	if err := repo.getMongoService(); err != nil {
		return nil, err
	}

	filter := bson.M{
		"product_id": cp.ProductId,
	}

	res := collectionmodels.Product{}

	if err := repo.DB.FindFirstData(repo.getCollection(), &res, filter); err != nil {
		return nil, err
	}

	return &res, nil
}

func (repo ProductRepository) ReadInProductId(productIds []string) ([]collectionmodels.Product, error) {

	if err := repo.getMongoService(); err != nil {
		return nil, err
	}

	inFilter := bson.A{}
	for _, v := range productIds {
		inFilter = append(inFilter, v)
	}
	filter := bson.M{
		"product_id": bson.M{"$in": inFilter},
	}

	res := make([]collectionmodels.Product, 0)

	if err := repo.DB.FindFirstData(repo.getCollection(), &res, filter); err != nil {
		return nil, err
	}

	return res, nil
}

func (repo ProductRepository) Update(cp collectionmodels.Product) (*models.UpdateProductO, error) {

	if err := repo.getMongoService(); err != nil {
		return nil, err
	}

	res, err := repo.DB.UpdateByID(repo.getCollection(), cp.Id, cp)
	if err != nil {
		return nil, err
	}

	if res.MatchedCount == 0 {
		return nil, fmt.Errorf("not found product at id [%s]", cp.Id.Hex())
	}

	return &models.UpdateProductO{}, nil
}

func (repo ProductRepository) Delete(cp collectionmodels.Product) error {

	if err := repo.getMongoService(); err != nil {
		return err
	}

	filter := bson.M{
		"_id": cp.Id,
	}

	res, err := repo.DB.DeleteOne(repo.getCollection(), filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("not found product at id [%s]", cp.Id.Hex())
	}

	return nil
}

func (repo ProductRepository) UpdateStock(cp collectionmodels.Product) (*models.ProductUpdateStockO, error) {

	if err := repo.getMongoService(); err != nil {
		return nil, err
	}

	res, err := repo.DB.UpdateByID(repo.getCollection(), cp.Id, cp)
	if err != nil {
		return nil, err
	}

	if res.MatchedCount == 0 {
		return nil, fmt.Errorf("not found product at id [%s]", cp.Id.Hex())
	}

	return &models.ProductUpdateStockO{}, nil
}
