package repository

import (
	"authservice/internal"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName         string
	Username         string
	Email            string
	Password         string
	PhotoProfile     string
	VerifiedEmail    bool
	RegistrationType string
}

type Role struct {
	gorm.Model
	NameRole string
}

func CoreToModel(input internal.User) User {
	return User{
		FullName:         input.FullName,
		Username:         input.Username,
		Email:            input.Email,
		Password:         input.Password,
		PhotoProfile:     input.PhotoProfile,
		VerifiedEmail:    input.VerifiedEmail,
		RegistrationType: input.RegistrationType,
	}
}

func CoreToModelUpdate(input internal.User) User {
	return User{
		FullName:     input.FullName,
		Username:     input.Username,
		Email:        input.Email,
		PhotoProfile: input.PhotoProfile,
	}
}

func (u User) ModelToCore() internal.User {
	return internal.User{
		ID:               u.ID,
		FullName:         u.FullName,
		Email:            u.Email,
		Password:         u.Password,
		PhotoProfile:     u.PhotoProfile,
		VerifiedEmail:    u.VerifiedEmail,
		RegistrationType: u.RegistrationType,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
	}
}
