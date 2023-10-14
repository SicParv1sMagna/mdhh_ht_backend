package repository

import (
	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
)

func (r *Repository) DeleteTalon(id int) error {
	err := r.db.Table("Talon").Delete(&model.Talon{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddTalon(talon *model.Talon) error {
	err := r.db.Table("Talon").Create(talon).Error
	if err != nil {
		return err
	}

	return nil
}
