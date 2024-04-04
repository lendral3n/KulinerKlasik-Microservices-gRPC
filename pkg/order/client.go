package order

import (
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/config"
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/order/pb"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect:", err)
	}
	return pb.NewOrderServiceClient(cc)
}