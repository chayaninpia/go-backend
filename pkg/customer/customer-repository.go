package customer

import (
	"apps/models"
	"apps/models/collectionmodels"
	"apps/pkg/utils/mongox"
	"apps/pkg/utils/tox"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"xorm.io/xorm"
)

type ICustomerRepository interface {
	// Create(tc tablemodels.TCustomer) error
	// Read(tc tablemodels.TCustomer) ([]models.CustomerReadO, error)
	// Update(tc tablemodels.TCustomer) (*models.CustomerUpdateO, error)
	// Delete(tc tablemodels.TCustomer) error

	Create(cc collectionmodels.Customer) error
	Read(cc collectionmodels.Customer) ([]models.CustomerReadO, error)
	Update(cc collectionmodels.Customer) (*models.CustomerUpdateO, error)
	Delete(cc collectionmodels.Customer) error
}

type CustomerRepository struct {
	E  *xorm.Engine
	DB *mongox.DBService
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

// func (repo *CustomerRepository) initDBService() error {
// 	if repo.E == nil {
// 		e, err := xormx.Init()
// 		if err != nil {
// 			return err
// 		}
// 		repo.E = e
// 	}
// 	return nil
// }

// mongodb
func (repo CustomerRepository) getCollection() *mongo.Collection {
	return repo.DB.Database.Collection(collectionmodels.Customer{}.CollectionName())
}

func (repo *CustomerRepository) getMongoService() error {

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

func (repo CustomerRepository) Create(cc collectionmodels.Customer) error {

	if err := repo.getMongoService(); err != nil {
		return err
	}

	if _, err := repo.DB.InsertOne(repo.getCollection(), cc); err != nil {
		return err
	}

	// if err := repo.initDBService(); err != nil {
	// 	return err
	// }

	// if _, err := repo.E.Cols(tc.Columns()...).Insert(tc); err != nil {
	// 	return err
	// }

	return nil
}

func (repo CustomerRepository) Read(cc collectionmodels.Customer) ([]models.CustomerReadO, error) {

	if err := repo.getMongoService(); err != nil {
		return nil, err
	}

	if len(tox.String(cc.ContactNo)) > 0 {
		return repo.readByContactNo(cc)
	} else if len(tox.String(cc.FullName)) > 0 {
		return repo.readByFullName(cc)
	}

	// if err := repo.initDBService(); err != nil {
	// 	return nil, err
	// }

	// qs := repo.E.Table(tablemodels.TCustomer{}.TableName())
	// if len(tc.ContactNo) > 0 {
	// 	return repo.readByContactNo(qs, tc)
	// } else if len(tc.FullName) > 0 {
	// 	return repo.readByFullName(qs, tc)
	// }

	return nil, nil
}

func (repo CustomerRepository) readByContactNo(cc collectionmodels.Customer) ([]models.CustomerReadO, error) {

	res := make([]collectionmodels.Customer, 0)
	filter := bson.M{
		"contact_no": bson.M{"$regex": cc.ContactNo},
	}

	if err := repo.DB.FindDatas(repo.getCollection(), &res, filter); err != nil {
		return nil, err
	}

	result := make([]models.CustomerReadO, 0)
	for _, v := range res {
		result = append(result, models.CustomerReadO{
			Id:        v.Id.Hex(),
			FullName:  tox.String(v.FullName),
			ContactNo: tox.String(v.ContactNo),
			Email:     tox.String(v.Email),
			Address:   tox.String(v.Address),
			Tambon:    tox.String(v.Tambon),
			Amphure:   tox.String(v.Amphure),
			Province:  tox.String(v.Province),
			PostCode:  tox.Int(v.PostCode),
		})
	}

	return result, nil

	// whereLike := fmt.Sprintf("%v%v", tc.ContactNo, "%")
	// qs.Where(`contact_no LIKE ?`, whereLike)
	// res, err := qs.QueryString()
	// if err != nil {
	// 	return nil, err
	// }

	// result := make([]models.CustomerReadO, 0)
	// for _, v := range res {
	// 	result = append(result, models.CustomerReadO{
	// 		Id:          v[`id`],
	// 		FullName:    v[`full_name`],
	// 		ContactNo:   v[`contact_no`],
	// 		Email:       v[`email`],
	// 		Address:     v[`address`],
	// 		SubDistrict: v[`sub_district`],
	// 		District:    v[`district`],
	// 		Province:    v[`province`],
	// 		PostCode:    v[`post_code`],
	// 	})
	// }

	// return result, nil
}

func (repo CustomerRepository) readByFullName(cc collectionmodels.Customer) ([]models.CustomerReadO, error) {

	res := make([]collectionmodels.Customer, 0)
	filter := bson.M{
		"full_name": bson.M{"$regex": cc.FullName},
	}

	if err := repo.DB.FindDatas(repo.getCollection(), &res, filter); err != nil {
		return nil, err
	}

	result := make([]models.CustomerReadO, 0)
	for _, v := range res {
		result = append(result, models.CustomerReadO{
			Id:        v.Id.Hex(),
			FullName:  tox.String(v.FullName),
			ContactNo: tox.String(v.ContactNo),
			Email:     tox.String(v.Email),
			Address:   tox.String(v.Address),
			Tambon:    tox.String(v.Tambon),
			Amphure:   tox.String(v.Amphure),
			Province:  tox.String(v.Province),
			PostCode:  tox.Int(v.PostCode),
		})
	}

	return result, nil
	// whereLike := fmt.Sprintf("%v%v%v", "%", tc.FullName, "%")
	// qs.Where(`full_name LIKE ?`, whereLike)
	// res, err := qs.QueryString()
	// if err != nil {
	// 	return nil, err
	// }

	// result := make([]models.CustomerReadO, 0)
	// for _, v := range res {
	// 	result = append(result, models.CustomerReadO{
	// 		Id:          v[`id`],
	// 		FullName:    v[`full_name`],
	// 		ContactNo:   v[`contact_no`],
	// 		Email:       v[`email`],
	// 		Address:     v[`address`],
	// 		SubDistrict: v[`sub_district`],
	// 		District:    v[`district`],
	// 		Province:    v[`province`],
	// 		PostCode:    v[`post_code`],
	// 	})
	// }

	// return result, nil
}

func (repo CustomerRepository) Update(cc collectionmodels.Customer) (*models.CustomerUpdateO, error) {

	if err := repo.getMongoService(); err != nil {
		return nil, err
	}

	updateRes, err := repo.DB.UpdateByID(repo.getCollection(), cc.Id, cc)
	if err != nil {
		return nil, err
	}

	if updateRes.MatchedCount == 0 {
		return nil, fmt.Errorf("not found customer at id [%s]", cc.Id.Hex())
	}

	return &models.CustomerUpdateO{
		Id:        cc.Id.Hex(),
		FullName:  tox.String(cc.FullName),
		ContactNo: tox.String(cc.ContactNo),
		Email:     tox.String(cc.Email),
		Address:   tox.String(cc.Address),
		Tambon:    tox.String(cc.Tambon),
		Amphure:   tox.String(cc.Amphure),
		Province:  tox.String(cc.Province),
		PostCode:  tox.Int(cc.PostCode),
	}, nil

	// if err := repo.initDBService(); err != nil {
	// 	return nil, err
	// }

	// updateCond := tablemodels.TCustomer{
	// 	Id: tc.Id,
	// }
	// rowAffect, err := repo.E.Update(tc, updateCond)
	// if err != nil {
	// 	return nil, nil
	// }

	// if rowAffect == 0 {
	// 	return nil, fmt.Errorf("not found customer")
	// }

	// return &models.CustomerUpdateO{
	// 	Id:          tc.Id,
	// 	FullName:    tc.FullName,
	// 	ContactNo:   tc.ContactNo,
	// 	Email:       tc.Email,
	// 	Address:     tc.Address,
	// 	SubDistrict: tc.SubDistrict,
	// 	District:    tc.District,
	// 	Province:    tc.Province,
	// 	PostCode:    tc.PostCode,
	// }, nil
}

func (repo CustomerRepository) Delete(cc collectionmodels.Customer) error {

	if err := repo.getMongoService(); err != nil {
		return err
	}

	filter := bson.M{
		"_id": cc.Id,
	}

	deleteRes, err := repo.DB.DeleteOne(repo.getCollection(), filter)
	if err != nil {
		return err
	}

	if deleteRes.DeletedCount == 0 {
		return fmt.Errorf("not found customer at id [%s]", cc.Id.Hex())
	}

	return nil
	// if err := repo.initDBService(); err != nil {
	// 	return err
	// }

	// deleteCond := tablemodels.TCustomer{
	// 	Id: tc.Id,
	// }

	// rowAffect, err := repo.E.Delete(deleteCond)
	// if err != nil {
	// 	return nil
	// }

	// if rowAffect == 0 {
	// 	return fmt.Errorf("not found customer")
	// }

	// return nil
}
