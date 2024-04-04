package auth

import (
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/config"
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/auth/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
    routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
    routes.Login(ctx, svc.Client)
}