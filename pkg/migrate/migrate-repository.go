package migrate

import (
	"apps/models/tablemodels"
	"apps/pkg/utils/gormx"

	"gorm.io/gorm"
)

type IMigrateRepository interface {
	AutoMigrate() error
}

type MigrateRepository struct {
	DB *gorm.DB
}

func NewMigrateRepository() *MigrateRepository {
	return &MigrateRepository{}
}

func (repo *MigrateRepository) initDBService() error {
	if repo.DB == nil {
		db, err := gormx.Init()
		if err != nil {
			return err
		}
		repo.DB = db
	}
	return nil
}
func (repo MigrateRepository) AutoMigrate() error {

	if err := gormx.AutoMigrate(); err != nil {
		return err
	}
	if err := repo.initDBService(); err != nil {
		return err
	}

	if err := repo.DB.Debug().AutoMigrate(

		&tablemodels.TProductGroup{},
		&tablemodels.TProduct{},
		&tablemodels.TProductSub{},
		&tablemodels.TProductStock{},
		&tablemodels.TSaleChannel{},
		&tablemodels.TSaleOrder{},
		&tablemodels.TSaleOrderItem{},
	); err != nil {
		return err
	}
	return nil
}
