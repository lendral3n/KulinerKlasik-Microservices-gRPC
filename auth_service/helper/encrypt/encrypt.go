package encrypts

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type HashInterface interface {
	CheckPasswordHash(hashed string, input string) bool
	HashPassword(input string) (string, error)
}

type hash struct{}

func New() HashInterface {
	return &hash{}
}

func (h *hash) HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("HASH - terjadi kesalahan saat hash password, error", err.Error())
		return "", err
	}

	return string(result), nil
}

func (h *hash) CheckPasswordHash(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil //true means login success
}
