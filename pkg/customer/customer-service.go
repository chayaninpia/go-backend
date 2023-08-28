package customer

import (
	"apps/models"
	"apps/models/collectionmodels"
	"apps/pkg/utils/timex"
	"apps/pkg/utils/validx"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ICustomerService interface {
	Create(req models.CustomerCreateI) error
	Read(req models.CustomerReadI) ([]models.CustomerReadO, error)
	Update(req models.CustomerUpdateI) (*models.CustomerUpdateO, error)
	Delete(req models.CustomerDeleteI) error
}

type CustomerService struct {
	CustomerRepository ICustomerRepository
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		CustomerRepository: NewCustomerRepository(),
	}
}

func (service CustomerService) Create(req models.CustomerCreateI) error {

	if err := validx.Struct(req); err != nil {
		return err
	}

	timeNow := timex.TimeNowPtr()
	cc := collectionmodels.Customer{
		Id:        primitive.NewObjectID(),
		FullName:  &req.FullName,
		ContactNo: &req.ContactNo,
		Email:     &req.Email,
		Address:   &req.Address,
		Tambon:    &req.Tambon,
		Amphure:   &req.Amphure,
		Province:  &req.Province,
		PostCode:  &req.PostCode,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	return service.CustomerRepository.Create(cc)
}

func (service CustomerService) Read(req models.CustomerReadI) ([]models.CustomerReadO, error) {

	cc := collectionmodels.Customer{
		FullName:  &req.FullName,
		ContactNo: &req.ContactNo,
	}

	return service.CustomerRepository.Read(cc)

}

func (service CustomerService) Update(req models.CustomerUpdateI) (*models.CustomerUpdateO, error) {
	if err := validx.Struct(req); err != nil {
		return nil, err
	}

	id, _ := primitive.ObjectIDFromHex(req.Id)
	cc := collectionmodels.Customer{
		Id:        id,
		FullName:  &req.FullName,
		ContactNo: &req.ContactNo,
		Email:     &req.Email,
		Address:   &req.Address,
		Tambon:    &req.Tambon,
		Amphure:   &req.Amphure,
		Province:  &req.Province,
		PostCode:  &req.PostCode,
		UpdatedAt: timex.DateNowPtr(),
	}

	return service.CustomerRepository.Update(cc)
}

func (service CustomerService) Delete(req models.CustomerDeleteI) error {

	id, _ := primitive.ObjectIDFromHex(req.Id)
	cc := collectionmodels.Customer{
		Id: id,
	}

	return service.CustomerRepository.Delete(cc)
}
