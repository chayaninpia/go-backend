package gormx

import "apps/models/tablemodels"

func AutoMigrate() error {

	db, err := Init()
	if err != nil {
		return err
	}

	if err := db.Debug().AutoMigrate(

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
