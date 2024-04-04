package routes

import (
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/menu/pb"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FineOne(ctx *gin.Context, c pb.MenuServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}