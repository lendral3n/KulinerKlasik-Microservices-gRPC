package auth

import (
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/auth/pb"
	"lendral3n/KulinerKlasik-Microservices-gRPC/pkg/config"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	
	if err != nil {
		fmt.Println("could not connect:", err)
	}
	return pb.NewAuthServiceClient(cc)
}