package routes

import (
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/order/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
	MenuId   int64 `json:"menuId"`
	Quantity int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")
	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		MenuId: body.MenuId,
		Quantity: body.Quantity,
		UserId: userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}