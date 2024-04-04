package routes

import (
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/menu/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMenuRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateMenu(ctx *gin.Context, c pb.MenuServiceClient) {
	body := CreateMenuRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateMenu(context.Background(), &pb.CreateMenuRequest{
		Name: body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}