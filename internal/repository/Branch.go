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

func (r *Repository) GetBranchById(id int) (model.Branch, error) {
	var branch model.Branch

	err := r.db.Table("Branch").Where(`"id" = ?`, id).First(&branch).Error
	if err != nil {
		return branch, err
	}

	return branch, nil
}

func (r *Repository) UpdateBranchTalonCount(id int, count int) error {
	result := r.db.Table("Branch").Where("id = ?", id).Update("talonCount", count)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}
