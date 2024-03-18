package middleware

import (
	"context"
	"fmt"
	"authservice/app/config"
	"time"
	"strings"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

// Generate token jwt
func CreateTokenLogin(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.JWT_SECRET))
	return signedToken, err
}


// Extract token jwt for gRPC
func ExtractTokenUserId(ctx context.Context) (int, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, status.Errorf(codes.Unauthenticated, "no auth metadata found in request")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return 0, status.Errorf(codes.Unauthenticated, "authorization token not provided")
	}

	jwtToken := values[0]
jwtToken = strings.TrimPrefix(jwtToken, "Bearer ")

	tokenJWT, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return 0, status.Errorf(codes.Unauthenticated, "error parsing token: %v", err)
	}

	if claims, ok := tokenJWT.Claims.(jwt.MapClaims); ok && tokenJWT.Valid {
		userId, isValidUserId := claims["userId"].(float64)
		if !isValidUserId {
			return 0, status.Errorf(codes.Unauthenticated, "invalid user id in token")
		}
		return int(userId), nil
	}

	return 0, status.Errorf(codes.Unauthenticated, "invalid token")
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
