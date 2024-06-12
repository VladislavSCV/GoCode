package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/VladislavSCV/GoCode/api/grpc/gen/pb-go/com.user_data"
	// "github.com/VladislavSCV/GoCode/migrations"
)

// server is used to implement pb.UserDataMessageServiceServer.
type server struct {
    pb.UnimplementedUserDataMessageServiceServer
}

// Implementing the methods for the UserDataMessageServiceServer.
func (s *server) GetUserData(ctx context.Context, in *pb.GetUserDataRequest) (*pb.GetUserDataResponse, error) {
    return nil, nil
}

func (s *server) UpdateUserData(ctx context.Context, in *pb.UpdateUserDataRequest) (*pb.UpdateUserDataResponse, error) {
    return nil, nil
}

func (s *server) UpdateUserPassword(ctx context.Context, in *pb.UpdateUserPasswordRequest) (*pb.UpdateUserPasswordResponse, error) {
    return nil, nil
}

func (s *server) UpdateUserEmail(ctx context.Context, in *pb.UpdateUserEmailRequest) (*pb.UpdateUserEmailResponse, error) {
    return nil, nil
}

func (s *server) UpdateUserPhone(ctx context.Context, in *pb.UpdateUserPhoneRequest) (*pb.UpdateUserPhoneResponse, error) {
    return nil, nil
}

func (s *server) UpdateUserStatus(ctx context.Context, in *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusResponse, error) {
    return nil, nil
}

func (s *server) UpdateUserAvatar(ctx context.Context, in *pb.UpdateUserAvatarRequest) (*pb.UpdateUserAvatarResponse, error) {
    return nil, nil
}

func (s *server) UpdateUserName(ctx context.Context, in *pb.UpdateUserNameRequest) (*pb.UpdateUserNameResponse, error) {
    return nil, nil
}

func (s *server) UpdateUserDescription(ctx context.Context, in *pb.UpdateUserDescriptionRequest) (*pb.UpdateUserDescriptionResponse, error) {
    return nil, nil
}

func connectGRPC() {
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
