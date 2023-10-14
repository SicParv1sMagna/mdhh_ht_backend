package repository

import "github.com/SicParv1sMagna/mdhh_backend/internal/model"

func (r *Repository) GetAllBranches() ([]model.Branch, error) {
	var branches []model.Branch

	err := r.db.Table("Branch").Find(&branches).Error
	if err != nil {
		return branches, err
	}

	return branches, nil
}

func (r *Repository) GetBranchBySearch(search string) ([]model.Branch, error) {
	var branches []model.Branch

	err := r.db.Table("Branch").Where(`"metroStation" LIKE ? OR "salePointName" LIKE ?`, "%"+search+"%", "%"+search+"%").Find(&branches).Error
	if err != nil {
		return nil, err
	}

	return branches, nil
}
