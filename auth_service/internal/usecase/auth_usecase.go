package usecase

import (
	encrypts "authservice/helper/encrypt"
	"authservice/helper/middleware"
	"authservice/internal"
	"errors"
)

type authUseCase struct {
	authData    internal.AuthRepositoryInterface
	hashService encrypts.HashInterface
}

// dependency injection
func New(repo internal.AuthRepositoryInterface, hash encrypts.HashInterface) internal.AuthUseCaseInterface {
	return &authUseCase{
		authData:    repo,
		hashService: hash,
	}
}

// Register implements internal.AuthUseCaseInterface.
func (a *authUseCase) Register(input internal.User) error {
	if input.Email == "" {
		return errors.New("email")
	}

	if input.Password != "" {
		hashedPasword, err := a.hashService.HashPassword(input.Password)
		if err != nil {
			return err
		}
		input.Password = hashedPasword
	}

	err := a.authData.Insert(input)
	if err != nil {
		return err
	}

	return nil
}

// Login implements internal.AuthUseCaseInterface.
func (a *authUseCase) Login(email string, password string) (data *internal.User, token string, err error) {
	if email == "" {
		return nil, "", errors.New("emai")
	}
	if password == "" {
		return nil, "", errors.New("password")
	}

	user, err := a.authData.Login(email, password)
	if err != nil {
		return nil, "", err
	}

	isValid := a.hashService.CheckPasswordHash(user.Password, password)
	if !isValid {
		return nil, "", errors.New("passwrd")
	}
	
	token, errJwt := middleware.CreateTokenLogin(int(user.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}

	return data, token, nil
}

// GetProfile implements internal.AuthUseCaseInterface.
func (a *authUseCase) GetProfile(userId int) (*internal.User, error) {
	result, err := a.authData.SelectById(userId)
	return result, err
}

// ChangePassword implements internal.AuthUseCaseInterface.
func (a *authUseCase) ChangePassword(userId int, oldPassword string, newPassword string) error {
	if oldPassword == "" {
		return errors.New("please input current password")
	}

	if newPassword == "" {
		return errors.New("please input new password")
	}

	hashedNewPass, errHash := a.hashService.HashPassword(newPassword)
	if errHash != nil {
		return errors.New("error hash password")
	}

	err := a.authData.ChangePassword(userId, oldPassword, hashedNewPass)
	if err != nil {
		return err
	}

	return nil
}

// DeleteAccount implements internal.AuthUseCaseInterface.
func (a *authUseCase) DeleteAccount(userId int) error {
	if userId <= 0 {
		return errors.New("invalid id")
	}
	err := a.authData.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser implements internal.AuthUseCaseInterface.
func (a *authUseCase) UpdateUser(userId int, input internal.User) error {
	err := a.authData.Update(userId, input)
	if err != nil {
		return err
	}
	return nil
}
