package delivery

import (
	"authservice/internal"
	pb "authservice/internal/delivery/grpc"
	"context"
)

type AuthServiceServer struct {
	uc internal.AuthUseCaseInterface
	pb.UnimplementedAuthServiceServer
}

func New(usecase internal.AuthUseCaseInterface) *AuthServiceServer {
	return &AuthServiceServer{
		uc: usecase,
	}
}

func (s *AuthServiceServer) Register(ctx context.Context, req *pb.User) (*pb.Response, error) {
	user := internal.User{
		FullName:         req.FullName,
		Username:         req.Username,
		Password:         req.Password,
		Email:            req.Email,
		PhotoProfile:     req.PhotoProfile,
		VerifiedEmail:    req.VerifiedEmail,
		RegistrationType: req.RegistrationType,
	}

	err := s.uc.Register(user)
	if err != nil {
		return &pb.Response{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.Response{
		Success: true,
		Message: "User registered successfully",
	}, nil
}

func (s *AuthServiceServer) GetProfile(ctx context.Context, req *pb.Request) (*pb.User, error) {
	user, err := s.uc.GetProfile(int(req.UserId))
	if err != nil {
		return nil, err
	}

	// Konversi internal User ke protobuf User
	protoUser := &pb.User{
		Id:               uint32(user.ID),
		FullName:         user.FullName,
		Username:         user.Username,
		Email:            user.Email,
		PhotoProfile:     user.PhotoProfile,
		VerifiedEmail:    user.VerifiedEmail,
		RegistrationType: user.RegistrationType,
		// Isi field lainnya sesuai kebutuhan
	}

	return protoUser, nil
}

func (s *AuthServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.Response, error) {
	// Konversi protobuf User ke internal User
	user := internal.User{
		ID:               uint(req.User.Id),
		FullName:         req.User.FullName,
		Username:         req.User.Username,
		Password:         req.User.Password,
		Email:            req.User.Email,
		PhotoProfile:     req.User.PhotoProfile,
		VerifiedEmail:    req.User.VerifiedEmail,
		RegistrationType: req.User.RegistrationType,
	}

	err := s.uc.UpdateUser(int(req.UserId), user)
	if err != nil {
		return &pb.Response{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.Response{
		Success: true,
		Message: "User updated successfully",
	}, nil
}

func (s *AuthServiceServer) DeleteAccount(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	err := s.uc.DeleteAccount(int(req.UserId))
	if err != nil {
		return &pb.Response{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.Response{
		Success: true,
		Message: "Account deleted successfully",
	}, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, token, err := s.uc.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	response := &pb.LoginResponse{
		FullName: user.FullName, 
		Token:    token,        
	}

	return response, nil
}


func (s *AuthServiceServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.Response, error) {
	err := s.uc.ChangePassword(int(req.UserId), req.OldPassword, req.NewPassword)
	if err != nil {
		return &pb.Response{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.Response{
		Success: true,
		Message: "Password changed successfully",
	}, nil
}
