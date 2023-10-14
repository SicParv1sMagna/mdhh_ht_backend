package repository

import (
	"fmt"

	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
)

func (r *Repository) GetAllAtms() ([]model.Atms, error) {
	var atms []model.Atms

	sql := `SELECT * FROM Atms`

	// Execute the SQL query
	err := r.db.Raw(sql).Scan(&atms).Error
	if err != nil {
		return atms, err
	}

	return atms, err
}

func (r *Repository) GetAtmById(id int) (model.Atms, error) {
	var atm model.Atms

	err := r.db.Where(`id=?`, id).First(&atm).Error
	if err != nil {
		return atm, err
	}
	fmt.Println(atm)
	return atm, err
}

func (r *Repository) GetAtmByName(name string) ([]model.Atms, error) {
	var atm []model.Atms

	err := r.db.Where(`address LIKE ?`, "%"+name+"%").Find(&atm).Error
	if err != nil {
		return atm, err
	}

	return atm, nil
}
