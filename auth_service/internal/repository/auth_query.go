package repository

import (
	"authservice/internal"
	"errors"

	"gorm.io/gorm"
)

type authQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) internal.AuthRepositoryInterface {
	return &authQuery{
		db: db,
	}
}

// Insert implements internal.AuthRepositoryInterface.
func (u *authQuery) Insert(input internal.User) error {
	data := CoreToModel(input)
	tx := u.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Login implements internal.AuthRepositoryInterface.
func (u *authQuery) Login(email string, password string) (data *internal.User, err error) {
	var userData User
	tx := u.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return nil, errors.New(" Invalid email or password")
	}
	result := userData.ModelToCore()
	return &result, nil
}

// SelectById implements internal.AuthRepositoryInterface.
func (u *authQuery) SelectById(userId int) (*internal.User, error) {
	var userData User
	tx := u.db.First(&userData, userId)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := userData.ModelToCore()
	return &result, nil
}

// Update implements internal.AuthRepositoryInterface.
func (u *authQuery) Update(userId int, input internal.User) error {
	data := CoreToModelUpdate(input)
	tx := u.db.Model(&User{}).Where("id = ?", userId).Updates(data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// ChangePassword implements internal.AuthRepositoryInterface.
func (u *authQuery) ChangePassword(userId int, oldPassword string, newPassword string) error {
	var userData User
	userData.Password = newPassword
	tx := u.db.Model(&User{}).Where("id = ?", userId).Updates(&userData)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements internal.AuthRepositoryInterface.
func (u *authQuery) Delete(userId int) error {
	tx := u.db.Delete(&User{}, userId)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
