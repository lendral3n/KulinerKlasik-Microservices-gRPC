package order

import (
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/auth"
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/order/routes"
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}