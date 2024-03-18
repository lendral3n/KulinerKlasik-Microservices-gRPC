package internal

import "time"

type User struct {
	ID               uint
	FullName         string
	Username         string
	Password         string
	Email            string
	PhotoProfile     string
	VerifiedEmail    bool
	RegistrationType string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Role Role
}

type Role struct {
	NameRole  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthUseCaseInterface interface {
	Register(input User) error
	GetProfile(userId int) (*User, error)
	UpdateUser(userId int, input User) error
	DeleteAccount(userId int) error
	Login(email, password string) (data *User, token string, err error)
	ChangePassword(userId int, oldPassword, newPassword string) error
}

type AuthRepositoryInterface interface {
	Insert(input User) error
	SelectById(userId int) (*User, error)
	Update(userId int, input User) error
	Delete(userId int) error
	Login(email, password string) (data *User, err error)
	ChangePassword(userId int, oldPassword, newPassword string) error
}
