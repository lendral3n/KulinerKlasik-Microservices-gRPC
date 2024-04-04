package menu

import (
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/config"
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/menu/pb"

	"google.golang.org/grpc"
	"fmt"
)

type ServiceClient struct {
	Client pb.MenuServiceClient
}

func InitServiceClient(c *config.Config) pb.MenuServiceClient {
	cc, err := grpc.Dial(c.MenuSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect:", err)
	}
	return pb.NewMenuServiceClient(cc)
}