package address

import (
	"apps/models"
	"apps/models/collectionmodels"
	"apps/pkg/constants"
	"apps/pkg/utils/filex"
	"apps/pkg/utils/timex"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAddressService interface {
	Create(req models.AddressCreateI) error
	CreateFromJson(req []models.AddressCreateI) error
	Read(req models.AddressReadI) ([]models.AddressReadO, error)
	Update(req models.AddressUpdateI) (*models.AddressUpdateO, error)
	Delete(req models.AddressDeleteI) error
	DeleteAll() error
}

type AddressService struct {
	AddressRepository IAddressRepository
}

func NewAddressService() *AddressService {
	return &AddressService{
		AddressRepository: NewAddressRepository(),
	}
}

func (service *AddressService) Create(req models.AddressCreateI) error {
	timeNow := timex.TimeNowPtr()

	createData := collectionmodels.Address{
		Id:        primitive.NewObjectID(),
		PostCode:  &req.PostCode,
		NameTh:    &req.NameTh,
		NameEn:    &req.NameEn,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Amphure: &collectionmodels.BaseAmphureModel{
			NameTh: &req.Amphure.NameTh,
			NameEn: &req.Amphure.NameEn,
			Province: &collectionmodels.BaseModel{
				NameTh: &req.Amphure.Province.NameTh,
				NameEn: &req.Amphure.Province.NameEn,
			},
		},
	}

	fileData, err := filex.ReadJSONFileArray[collectionmodels.Address](constants.PathFileAddress)
	if err != nil {
		return err
	}

	fileData = append(fileData, createData)

	if err := filex.WriteFile(fileData, constants.PathFileAddress); err != nil {
		return err
	}

	return service.AddressRepository.Create(createData)
}

func (service *AddressService) CreateFromJson(req []models.AddressCreateI) error {
	timeNow := timex.TimeNowPtr()

	createData := make([]collectionmodels.Address, 0)
	for _, v := range req {
		createData = append(createData, collectionmodels.Address{
			Id:        primitive.NewObjectID(),
			PostCode:  &v.PostCode,
			NameTh:    &v.NameTh,
			NameEn:    &v.NameEn,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
			Amphure: &collectionmodels.BaseAmphureModel{
				NameTh: &v.Amphure.NameTh,
				NameEn: &v.Amphure.NameEn,
				Province: &collectionmodels.BaseModel{
					NameTh: &v.Amphure.Province.NameTh,
					NameEn: &v.Amphure.Province.NameEn,
				},
			},
		})
	}

	if err := filex.WriteFile(createData, constants.PathFileAddress); err != nil {
		return err
	}

	return service.AddressRepository.CreateMany(createData)
}

func (service *AddressService) Read(req models.AddressReadI) ([]models.AddressReadO, error) {
	return service.AddressRepository.Read(req)
}

func (service *AddressService) Update(req models.AddressUpdateI) (*models.AddressUpdateO, error) {
	return service.AddressRepository.Update(req)
}

func (service *AddressService) Delete(req models.AddressDeleteI) error {
	return service.AddressRepository.Delete(req)
}

func (service *AddressService) DeleteAll() error {
	return service.AddressRepository.DeleteAll()
}
