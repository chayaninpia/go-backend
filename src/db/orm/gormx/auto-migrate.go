package gormx

import (
	"apps/src/db/models/tb"
)

func AutoMigrate() error {

	db, err := Init()
	if err != nil {
		return err
	}

	if err := db.Debug().AutoMigrate(

		&tb.TProductGroup{},
		&tb.TProduct{},
		&tb.TProductSub{},
		&tb.TProductStock{},
		&tb.TSaleChannel{},
		&tb.TSaleOrder{},
		&tb.TSaleOrderItem{},
	); err != nil {
		return err
	}

	return nil
}
