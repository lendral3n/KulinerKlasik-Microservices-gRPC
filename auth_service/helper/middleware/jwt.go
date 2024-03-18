package middleware

import (
	"authservice/app/config"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.JWT_SECRET),
		SigningMethod: "HS256",
	})
}

// Generate token jwt
func CreateTokenLogin(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_SECRET))
}

// extract token jwt
func ExtractTokenUserId(e echo.Context) int {
	header := e.Request().Header.Get("Authorization")
	headerToken := strings.Split(header, " ")
	token := headerToken[len(headerToken)-1]
	tokenJWT, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if tokenJWT.Valid {
		claims := tokenJWT.Claims.(jwt.MapClaims)
		userId, isValidUserId := claims["userId"].(float64)
		if !isValidUserId {
			return 0
		}
		return int(userId)
	}
	return 0
}

func CreateResetPasswordToken(userId int) (string, error) {
	payload := map[string]interface{}{
		"userId":        userId,
		"resetPassword": true,
	}

	now := time.Now().UTC()

	claims := jwt.MapClaims{
		"sub": payload,
		"exp": now.Add(time.Duration(time.Minute * 15)).Unix(),
		"iat": now.Unix(),
		"nbf": now.Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func ExtractUserIdFromResetPasswordToken(token string) (int, error) {
	tokenJWT, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := tokenJWT.Claims.(jwt.MapClaims); ok && tokenJWT.Valid {
		payload, ok := claims["sub"].(map[string]interface{})
		if !ok {
			return 0, fmt.Errorf("invalid payload in token")
		}

		userId, isValidUserId := payload["userId"].(float64)
		if !isValidUserId {
			return 0, fmt.Errorf("invalid user id in token")
		}
		return int(userId), nil
	}

	return 0, fmt.Errorf("invalid token")
}
