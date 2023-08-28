package address

import (
	"apps/models"
	"apps/models/collectionmodels"
	"apps/pkg/constants"
	"apps/pkg/utils/filex"
	"apps/pkg/utils/mongox"
	"apps/pkg/utils/slicex"
	"apps/pkg/utils/timex"
	"apps/pkg/utils/tox"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAddressRepository interface {
	Create(cad collectionmodels.Address) error
	CreateMany(cad []collectionmodels.Address) error
	Read(adc models.AddressReadI) ([]models.AddressReadO, error)
	Update(adc models.AddressUpdateI) (*models.AddressUpdateO, error)
	Delete(adc models.AddressDeleteI) error
	DeleteAll() error
}

type AddressRepository struct {
	DB *mongox.DBService
}

func NewAddressRepository() *AddressRepository {
	return &AddressRepository{}
}

func (repo AddressRepository) getCollection() *mongo.Collection {
	return repo.DB.Database.Collection(collectionmodels.Address{}.CollectionName())
}

func (repo *AddressRepository) getMongoService() error {

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

func (repo AddressRepository) Create(cad collectionmodels.Address) error {

	if err := repo.getMongoService(); err != nil {
		return err
	}

	if _, err := repo.DB.InsertOne(repo.getCollection(), cad); err != nil {
		return err
	}
	return nil
}

func (repo AddressRepository) CreateMany(cad []collectionmodels.Address) error {

	if err := repo.getMongoService(); err != nil {
		return err
	}

	toInsert := make([]interface{}, len(cad))
	for index, value := range cad {
		toInsert[index] = value
	}

	if _, err := repo.DB.InsertMany(repo.getCollection(), toInsert); err != nil {
		return err
	}
	return nil
}

func (repo AddressRepository) Read(adc models.AddressReadI) ([]models.AddressReadO, error) {

	result, err := filex.ReadJSONFileArray[collectionmodels.Address](constants.PathFileAddress)
	if err != nil {
		return nil, err
	}

	res := make([]models.AddressReadO, 0)
	for _, v := range result {
		if strings.Contains(tox.String(v.NameTh), adc.Filter) || strings.Contains(tox.String(v.Amphure.NameTh), adc.Filter) || strings.Contains(tox.String(v.Amphure.Province.NameTh), adc.Filter) || len(adc.Filter) == 0 {
			res = append(res, models.AddressReadO{
				Id:       v.Id.Hex(),
				PostCode: tox.Int(v.PostCode),
				NameTh:   tox.String(v.NameTh),
				NameEn:   tox.String(v.NameEn),
				Amphure: models.BaseAmphureModel{
					NameTh: tox.String(v.Amphure.NameTh),
					NameEn: tox.String(v.Amphure.NameEn),
					Province: models.BaseProvinceModel{
						NameTh: tox.String(v.Amphure.Province.NameTh),
						NameEn: tox.String(v.Amphure.Province.NameEn),
					},
				},
			})
		}
	}

	return res, nil
}

func (repo AddressRepository) Update(adc models.AddressUpdateI) (*models.AddressUpdateO, error) {

	if err := repo.getMongoService(); err != nil {
		return nil, err
	}

	fileData, err := filex.ReadJSONFileArray[collectionmodels.Address](constants.PathFileAddress)
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(adc.Id)
	if err != nil {
		return nil, err
	}

	timeNow := timex.TimeNowPtr()
	toUpdate := collectionmodels.Address{
		Id:        id,
		PostCode:  &adc.PostCode,
		NameTh:    &adc.NameTh,
		NameEn:    &adc.NameEn,
		UpdatedAt: timeNow,
		Amphure: &collectionmodels.BaseAmphureModel{
			NameTh: &adc.Amphure.NameEn,
			NameEn: &adc.Amphure.NameTh,
			Province: &collectionmodels.BaseModel{
				NameTh: &adc.Amphure.Province.NameTh,
				NameEn: &adc.Amphure.Province.NameEn,
			},
		},
	}
	isUpdate := false
	for k, v := range fileData {
		if v.Id.Hex() == adc.Id {
			fileData[k] = toUpdate
			isUpdate = true
			break
		}
	}
	if !isUpdate {
		return nil, fmt.Errorf("not found address at id [%s]", id.Hex())
	}

	result := models.AddressUpdateO(adc)
	if _, err := repo.DB.UpdateByID(repo.getCollection(), id, toUpdate); err != nil {
		return nil, err
	}

	if err := filex.WriteFile(fileData, constants.PathFileAddress); err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo AddressRepository) Delete(adc models.AddressDeleteI) error {

	if err := repo.getMongoService(); err != nil {
		return err
	}

	fileData, err := filex.ReadJSONFileArray[collectionmodels.Address](constants.PathFileAddress)
	if err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(adc.Id)
	if err != nil {
		return err
	}

	keyDelete := -1
	for k, v := range fileData {
		if v.Id.Hex() == adc.Id {
			keyDelete = k
			break
		}
	}

	if keyDelete == -1 {
		return fmt.Errorf("not found address at id [%s]", id.Hex())
	}

	fileData = slicex.RemoveIndex[collectionmodels.Address](keyDelete, fileData)

	if err := filex.WriteFile(fileData, constants.PathFileAddress); err != nil {
		return err
	}

	filter := bson.M{"_id": id}

	if _, err := repo.DB.DeleteOne(repo.getCollection(), filter); err != nil {
		return err
	}

	return nil
}

func (repo AddressRepository) DeleteAll() error {

	if err := repo.getMongoService(); err != nil {
		return err
	}

	filter := bson.M{}

	if _, err := repo.DB.DeleteMany(repo.getCollection(), filter); err != nil {
		return err
	}

	if err := os.Remove(constants.PathFileAddress); err != nil {
		return err
	}

	return nil
}
