package menu

import (
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/auth"
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/menu/routes"
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc  := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/menu")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateMenu)
	routes.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
    routes.FineOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateMenu(ctx *gin.Context) {
    routes.CreateMenu(ctx, svc.Client)
}