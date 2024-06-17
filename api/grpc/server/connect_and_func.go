package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/VladislavSCV/GoCode/api/grpc/gen/pb-go/com.user_data"
	"github.com/VladislavSCV/GoCode/migrations"
)

func init() {
	migrations.ConnectToDB()
}

// server is used to implement pb.UserDataMessageServiceServer.
type server struct {
	pb.UnimplementedUserDataMessageServiceServer
}


// Login checks if the user is valid and returns a boolean indicating if the
// login was accepted and an error if there was one.
func (s *server) Login(ctx context.Context, in *pb.LoginUserDataRequest) (*pb.IsAcceptedLoginResponse, error) {
	isAccepted, err := migrations.Login(migrations.DB, in.Email, in.PasswordHash)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to login: %v", err)
	}
	return &pb.IsAcceptedLoginResponse{IsAccepted: isAccepted}, nil
}

// SignUp creates a new user in the database and returns a boolean indicating if
// the sign up was accepted and an error if there was one.
func (s *server) SignUp(ctx context.Context, in *pb.SignUserDataRequest) (*pb.SignUserDataResponse, error) {
	dateOfBirth, err := time.Parse(time.RFC3339, in.DateOfBirth.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to parse date of birth: %v", err)
	}
	err = migrations.SignUp(migrations.DB, in.Username, in.PasswordHash, in.Email, in.Phone, in.AvatarUrl, in.Status, in.Role, dateOfBirth)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to sign up: %v", err)
	}
	return &pb.SignUserDataResponse{IsAccepted: true}, nil
}



func (s *server) GetUserData(
	ctx context.Context,
	in *pb.GetUserDataRequest,
) (*pb.GetUserDataResponse, error) {
	userId := in.GetUserId()
	userData, err := migrations.GetUserData(migrations.DB, int(userId))
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Failed to fetch user data: %v",
			err,
		)
	}

	response := &pb.GetUserDataResponse{
		UserId:       int64(userData.ID),
		Username:     userData.UserName,
		Description:  userData.Description,
		Email:        userData.Email,
		Phone:        userData.Phone,
		AvatarUrl:    userData.AvatarUrl,
		Status:       userData.Status,
		Role:         userData.Role,
		PasswordHash: userData.PasswordHash,
		DateOfBirth:  userData.DateOfBirth.Format(time.RFC3339),
		PrivacySettings: userData.PrivacySettings,
		IsActive:     userData.IsActive,
		LastLogin:    userData.LastLogin.Format(time.RFC3339),
		ConfirmationToken: userData.ConfirmationToken,
		SocialProfiles:    userData.SocialProfiles,
		CreatedAt:         userData.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         userData.UpdatedAt.Format(time.RFC3339),
	}

	return response, nil
}


func (s *server) UpdateUserData(
	ctx context.Context,
	in *pb.UpdateUserDataRequest,
) (*pb.UpdateUserDataResponse, error) {
	dob, _ := time.Parse(time.RFC3339, in.DateOfBirth)
	ll, _ := time.Parse(time.RFC3339, in.LastLogin)
	ca, _ := time.Parse(time.RFC3339, in.CreatedAt)
	ua, _ := time.Parse(time.RFC3339, in.UpdatedAt)
	userData := &migrations.UserData{
		ID:               int(in.GetUserId()),
		UserName:         in.GetUsername(),
		Description:      in.GetDescription(),
		Email:            in.GetEmail(),
		Phone:            in.GetPhone(),
		AvatarUrl:        in.GetAvatarUrl(),
		Status:           in.GetStatus(),
		Role:             in.GetRole(),
		PasswordHash:     in.GetPasswordHash(),
		DateOfBirth:      dob,
		PrivacySettings:  in.GetPrivacySettings(),
		IsActive:         in.GetIsActive(),
		LastLogin:        ll,
		ConfirmationToken: in.GetConfirmationToken(),
		SocialProfiles:    in.GetSocialProfiles(),
		CreatedAt:         ca,
		UpdatedAt:         ua,
	}

	err := migrations.UpdateUserData(migrations.DB, userData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user data: %v", err)
	}

	return &pb.UpdateUserDataResponse{
		Success: true,
	}, nil
}

func (s *server) UpdateUserPassword(ctx context.Context, in *pb.UpdateUserPasswordRequest) (*pb.UpdateUserPasswordResponse, error) {
	err := migrations.UpdateUserPassword(migrations.DB, int(in.GetUserId()), in.GetNewPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user password: %v", err)
	}
	return &pb.UpdateUserPasswordResponse{
		Success: true,
	}, nil
}

func (s *server) UpdateUserEmail(ctx context.Context, in *pb.UpdateUserEmailRequest) (*pb.UpdateUserEmailResponse, error) {
	err := migrations.UpdateUserEmail(migrations.DB, int(in.GetUserId()), in.GetNewEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user email: %v", err)
	}
	return &pb.UpdateUserEmailResponse{
		Success: true,
	}, nil
}

func (s *server) UpdateUserPhone(ctx context.Context, in *pb.UpdateUserPhoneRequest) (*pb.UpdateUserPhoneResponse, error) {
	err := migrations.UpdateUserPhone(migrations.DB, int(in.GetUserId()), in.GetNewPhone())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user phone: %v", err)
	}
	return &pb.UpdateUserPhoneResponse{
		Success: true,
	}, nil
}

func (s *server) UpdateUserStatus(ctx context.Context, in *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusResponse, error) {
	err := migrations.UpdateUserStatus(migrations.DB, int(in.GetUserId()), in.GetNewStatus())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user status: %v", err)
	}
	return &pb.UpdateUserStatusResponse{
		Success: true,
	}, nil
}

func (s *server) UpdateUserAvatar(ctx context.Context, in *pb.UpdateUserAvatarRequest) (*pb.UpdateUserAvatarResponse, error) {
	err := migrations.UpdateUserAvatar(migrations.DB, int(in.GetUserId()), in.GetNewAvatarUrl())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user avatar: %v", err)
	}
	return &pb.UpdateUserAvatarResponse{
		Success: true,
	}, nil
}

func (s *server) UpdateUserName(ctx context.Context, in *pb.UpdateUserNameRequest,) (*pb.UpdateUserNameResponse, error) {
	err := migrations.UpdateUserName(migrations.DB, int(in.GetUserId()), in.GetNewUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user name: %v", err)
	}
	return &pb.UpdateUserNameResponse{
		Success: true,
	}, nil
}

func (s *server) UpdateUserDescription(ctx context.Context, in *pb.UpdateUserDescriptionRequest) (*pb.UpdateUserDescriptionResponse, error) {
	err := migrations.UpdateUserDescription(migrations.DB, int(in.GetUserId()), in.GetNewDescription())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user description: %v", err)
	}
	return &pb.UpdateUserDescriptionResponse{
		Success: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Starting gRPC server on port 50051...")

	s := grpc.NewServer()
	pb.RegisterUserDataMessageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
