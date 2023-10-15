package repository

import (
	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
)

func (r *Repository) AddRoute(route *model.Route) error {
	err := r.db.Table("History").Create(route).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAllRoutes() ([]model.Route, error) {
	var routes []model.Route

	err := r.db.Table("History").Find(&routes).Error
	if err != nil {
		return routes, err
	}

	return routes, nil
}
