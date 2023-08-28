package salechannel

import (
	"apps/models"
	"apps/models/tablemodels"
	"apps/pkg/utils/xormx"

	"github.com/google/uuid"
	"xorm.io/xorm"
)

type ISaleChannelRepository interface {
	Create(csc models.CreateSaleChannelI) error
	Delete(dsc models.DeleteSaleChannelI) error
	Read(rsc models.ReadSaleChannelI) ([]models.ReadSaleChannelO, error)
	Update(usc models.UpdateSaleChannelI) error
}

type SaleChannelRepository struct {
	E *xorm.Engine
}

func NewSaleChannelRepository() *SaleChannelRepository {
	return &SaleChannelRepository{}
}

func (repo *SaleChannelRepository) initDBService() error {
	if repo.E == nil {
		e, err := xormx.Init()
		if err != nil {
			return err
		}
		repo.E = e
	}
	return nil
}

func (repo SaleChannelRepository) Create(csc models.CreateSaleChannelI) error {

	if err := repo.initDBService(); err != nil {
		return err
	}

	saleChannelCreate := tablemodels.TSaleChannel{
		Id:          uuid.NewString(),
		SaleChannel: csc.SaleChannel,
	}

	if _, err := repo.E.Cols(saleChannelCreate.Columns()...).Insert(saleChannelCreate); err != nil {
		return nil
	}

	return nil
}

func (repo SaleChannelRepository) Delete(dsc models.DeleteSaleChannelI) error {

	if err := repo.initDBService(); err != nil {
		return err
	}

	saleChannelDelete := tablemodels.TSaleChannel{
		Id: dsc.Id,
	}

	if _, err := repo.E.Delete(saleChannelDelete); err != nil {
		return err
	}

	return nil
}

func (repo SaleChannelRepository) Read(rsc models.ReadSaleChannelI) ([]models.ReadSaleChannelO, error) {

	if err := repo.initDBService(); err != nil {
		return nil, err
	}

	qs := repo.E.Select(`*`).Table(tablemodels.TSaleChannel{}.TableName())

	if rsc.SaleChannel != `` {
		qs.Where(`sale_channel like ?`, `%`+rsc.SaleChannel+`%`)
	}

	result, err := qs.QueryString()
	if err != nil {
		return nil, err
	}

	res := make([]models.ReadSaleChannelO, 0)
	for _, v := range result {
		res = append(res, models.ReadSaleChannelO{
			Id:          v[`id`],
			SaleChannel: v[`sale_channel`],
		})
	}

	return res, nil

}

func (repo SaleChannelRepository) Update(usc models.UpdateSaleChannelI) error {

	if err := repo.initDBService(); err != nil {
		return err
	}

	saleChannelUpdate := tablemodels.TSaleChannel{
		SaleChannel: usc.SaleChannel,
	}
	saleChannelUpdateCondi := tablemodels.TSaleChannel{
		Id: usc.Id,
	}

	if _, err := repo.E.Update(saleChannelUpdate, saleChannelUpdateCondi); err != nil {
		return err
	}

	return nil
}
