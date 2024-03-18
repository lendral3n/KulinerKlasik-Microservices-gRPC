package delivery

import (
	"authservice/helper/middleware"
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

func (s *AuthServiceServer) Register(ctx context.Context, req *pb.UserRequest) (*pb.SuccesResponse, error) {
	user := internal.User{
		FullName:     req.FullName,
		Username:     req.Username,
		Email:        req.Email,
		Password:     req.Password,
		PhotoProfile: req.PhotoProfile,
	}

	err := s.uc.Register(user)
	if err != nil {
		return &pb.SuccesResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.SuccesResponse{
		Success: true,
		Message: "User registered successfully",
	}, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, token, err := s.uc.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	response := &pb.LoginResponse{
		FullName: user.FullName,
		JwtToken: token,
	}

	return response, nil
}

func (s *AuthServiceServer) GetProfile(ctx context.Context, req *pb.Empty) (*pb.UserGetProfile, error) {
	userId, err := middleware.ExtractTokenUserId(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.uc.GetProfile(int(userId))
	if err != nil {
		return nil, err
	}

	// Konversi internal User ke protobuf User
	protoUser := &pb.UserGetProfile{
		FullName:         user.FullName,
		Username:         user.Username,
		Email:            user.Email,
		PhotoProfile:     user.PhotoProfile,
		VerifiedEmail:    user.VerifiedEmail,
		RegistrationType: user.RegistrationType,
		Role:             user.Role.NameRole,
	}

	return protoUser, nil
}

func (s *AuthServiceServer) UpdateUser(ctx context.Context, req *pb.UserRequest) (*pb.SuccesResponse, error) {
	userId, err := middleware.ExtractTokenUserId(ctx)
	if err != nil {
		return nil, err
	}

	// Konversi protobuf User ke internal User
	user := internal.User{
		FullName:     req.FullName,
		Username:     req.Username,
		Password:     req.Password,
		Email:        req.Email,
		PhotoProfile: req.PhotoProfile,
	}

	err = s.uc.UpdateUser(userId, user)
	if err != nil {
		return &pb.SuccesResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.SuccesResponse{
		Success: true,
		Message: "User updated successfully",
	}, nil
}

func (s *AuthServiceServer) DeleteAccount(ctx context.Context, req *pb.Empty) (*pb.SuccesResponse, error) {
	
	userId, err := middleware.ExtractTokenUserId(ctx)
	if err != nil {
		return nil, err
	}

	err = s.uc.DeleteAccount(userId)
	if err != nil {
		return &pb.SuccesResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.SuccesResponse{
		Success: true,
		Message: "Account deleted successfully",
	}, nil
}

func (s *AuthServiceServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.SuccesResponse, error) {
	userId, err := middleware.ExtractTokenUserId(ctx)
	if err != nil {
		return nil, err
	}

	err = s.uc.ChangePassword(int(userId), req.OldPassword, req.NewPassword)
	if err != nil {
		return &pb.SuccesResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.SuccesResponse{
		Success: true,
		Message: "Password changed successfully",
	}, nil
}
