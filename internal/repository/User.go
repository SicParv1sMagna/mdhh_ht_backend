package repository

import "github.com/SicParv1sMagna/mdhh_backend/internal/model"

func (r *Repository) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	err := r.db.Table(`"User"`).Where("Email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) CreateUser(user model.User) error {
	err := r.db.Table(`"User"`).Create(&user).Error
	return err
}
