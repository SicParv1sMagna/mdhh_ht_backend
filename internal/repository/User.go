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
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserByToken(code string) (model.User, error) {
	var user model.User

	err := r.db.Table("User").Where("accesstoken = ?", code).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}

func (r *Repository) ConfirmRegistration(email string) error {
	err := r.db.Table("User").Where("Email = ?", email).Update("is_confirmed", true).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateUserAccessToken(email, code string) error {
	err := r.db.Table("User").Where("Email = ?", email).Update("accesstoken", code).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserById(id int) (model.User, error) {
	var user model.User

	err := r.db.Table("User").Where("User_ID=?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) UpdateUserData(user *model.User) error {
	err := r.db.Table("User").Update("Email, FirstName, SeondName, LegalEntity", &user).Error

	return err
}

func (r *Repository) DeleteUser(id int) error {
	err := r.db.Table("User").Where("User_ID = ?", id).Delete(&model.User{}).Error

	return err
}
