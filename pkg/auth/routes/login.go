package routes

import (
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/auth/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	body := LoginRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req, err := c.Login(context.Background(), &pb.LoginRequest{
		Email: body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &req)
}