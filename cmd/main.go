package main

import (
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/auth"
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/config"

	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/order"

	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/menu"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	menu.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)
	r.Run(c.Port)
}
