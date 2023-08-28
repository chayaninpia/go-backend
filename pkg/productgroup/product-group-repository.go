package productgroup

import (
	"apps/models"
	"apps/models/tablemodels"
	"apps/pkg/utils/xormx"
	"fmt"

	"github.com/google/uuid"
	"xorm.io/xorm"
)

type IProductGroupRepository interface {
	Create(cp models.CreateProductGroupI) error
	Delete(dp models.DeleteProductGroupI) error
	Read(rp models.ReadProductGroupI) ([]tablemodels.TProductGroup, error)
	Update(up models.UpdateProductGroupI) error
}

type ProductGroupRepository struct {
	E *xorm.Engine
}

func NewProductGroupRepository() *ProductGroupRepository {
	return &ProductGroupRepository{}
}

func (repo *ProductGroupRepository) initDBService() error {
	if repo.E == nil {
		e, err := xormx.Init()
		if err != nil {
			return err
		}
		repo.E = e
	}
	return nil
}

func (repo ProductGroupRepository) Create(cpg models.CreateProductGroupI) error {

	if err := repo.initDBService(); err != nil {
		return err
	}

	pdgInsert := tablemodels.TProductGroup{
		Id:        uuid.NewString(),
		GroupName: cpg.GroupName,
	}

	if _, err := repo.E.Cols(pdgInsert.Columns()...).Insert(pdgInsert); err != nil {
		return err
	}

	return nil
}

func (repo ProductGroupRepository) Delete(dpg models.DeleteProductGroupI) error {

	if err := repo.initDBService(); err != nil {
		return err
	}

	productGroupDelete := tablemodels.TProductGroup{
		Id: dpg.Id,
	}

	if _, err := repo.E.Delete(productGroupDelete); err != nil {
		return err
	}

	return nil
}

func (repo ProductGroupRepository) Read(rp models.ReadProductGroupI) ([]tablemodels.TProductGroup, error) {

	if err := repo.initDBService(); err != nil {
		return nil, err
	}

	res := make([]tablemodels.TProductGroup, 0)

	qs := repo.E.Select(`*`).Table(tablemodels.TProductGroup{}.TableName())

	if rp.GroupName != `` {
		whereLike := fmt.Sprintf("%v%v%v", "%", rp.GroupName, "%")
		qs.Where(`group_name LIKE ?`, whereLike)
	}

	if err := qs.Find(&res); err != nil {
		return nil, err
	}

	return res, nil

}

func (repo ProductGroupRepository) Update(upg models.UpdateProductGroupI) error {

	if err := repo.initDBService(); err != nil {
		return err
	}

	productGroupUpdate := tablemodels.TProductGroup{
		GroupName: upg.GroupName,
	}
	productGroupUpdateCondi := tablemodels.TProductGroup{
		Id: upg.Id,
	}

	if _, err := repo.E.Update(productGroupUpdate, productGroupUpdateCondi); err != nil {
		return err
	}

	return nil
}
