# KulinerKlasik-Microservices-gRPC

![KulinerKlasik-Microservices-gRPC](/pkg/docs/photo_6088876605150705876_y.jpg)
We are building a microservice architecture using Golang and gRPC. The architecture consists of three services: User Service, Menu Service, and Order Service. The User Service handles authentication and registration, the Menu Service manages product creation and retrieval, and the Order Service handles order management. To streamline communication between services, we are implementing an API Gateway. The API Gateway acts as a central entry point, routing client requests to the appropriate microservices. This architecture allows for independent development, deployment, and scaling of services, while providing a unified interface for clients.

## Dependency repositories

- **Service Auth** ==> https://github.com/lendral3n/KulinerKlasik-Microservices-gRPC-Service-Auth
- **Service Menu** ==> https://github.com/lendral3n/KulinerKlasik-Microservices-gRPC-Service-Menu
- **Service Order** ==> https://github.com/lendral3n/KulinerKlasik-Microservices-gRPC-Service-Order

## Application Infrastructure

1. **API Gateway:** Handles incoming HTTP requests
2. **Auth Service:** Provides features such as Register, Login and generates Token by JWT
3. **Menu Service:** Provides features such as Add Product, Decrease Stock and Find Product
4. **Order Service:** The only feature we ship in this Microservice is Create Order

## How to start ?

- `STEP 1 : `We need to clone [Service Auth](https://github.com/lendral3n/KulinerKlasik-Microservices-gRPC-Service-Auth)
- `STEP 2 : `Setup the environment in `local.env` like `local.env.example` according to the settings on your localhost
- `STEP 3 : `Run command on project's terminal : `go mod tidy` and `go mod vendor` to constructs a directory named vendor in the main module's root directory that contains copies of all packages needed to support builds and tests of packages in the main module
- `STEP 4 : `Run command `Make Server` or `go run cmd/main.go` as usualy
- `STEP 5 : `Clone [Service Menu](https://github.com/lendral3n/KulinerKlasik-Microservices-gRPC-Service-Menu)
- `STEP 6 : `Repeat `STEP 2` until `STEP 4`
- `STEP 7 : `Clone [Service Order](https://github.com/lendral3n/KulinerKlasik-Microservices-gRPC-Service-Order)
- `STEP 8 : `Repeat `STEP 2` until `STEP 4`
- `STEP 9 : `Dont forget to do `STEP 2` until `STEP 3` in this [KulinerKlasik Services](https://github.com/lendral3n/KulinerKlasik-Microservices-gRPC)
- `STEP 10 : `Execute command `Make Server` or `go run cmd/main.go` (again) in this [KulinerKlasik Services](https://github.com/lendral3n/KulinerKlasik-Microservices-gRPC)
