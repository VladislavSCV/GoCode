package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"

    pb "github.com/VladislavSCV/api/grpc/gen/pb-go/com.user_data"

    // "github.com/VladislavSCV/migrations/..." // импортируем все функции и структуры из пакета migrations
)


type server struct {
    pb.UnimplementedUserDataMessageServiceServer
}

func (s *server) GetUserData(ctx context.Context, in *pb.GetUserDataRequest) (*pb.GetUserDataResponse, error) {
    
}

func (s *server) UpdateUserData(ctx context.Context, in *pb.UpdateUserDataRequest) (*pb.UpdateUserDataResponse, error) {
    
}

func (s *server) UpdateUserPassword(ctx context.Context, in *pb.UpdateUserPasswordRequest) (*pb.UpdateUserPasswordResponse, error) {
    
}
func (s *server) UpdateUserEmail(ctx context.Context, in *pb.UpdateUserEmailRequest) (*pb.UpdateUserEmailResponse, error) {
    
}
func (s *server) UpdateUserPhone(ctx context.Context, in *pb.UpdateUserPhoneRequest) (*pb.UpdateUserPhoneResponse, error) {
    
}
func (s *server) UpdateUserStatus(ctx context.Context, in *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusResponse, error) {
    
}
func (s *server) UpdateUserAvatar(ctx context.Context, in *pb.UpdateUserDataRequest) (*pb.UpdateUserAvatarResponse, error) {
    
}

func (s *server) UpdateUserName(ctx context.Context, in *pb.UpdateUserNameRequest) (*pb.UpdateUserNameResponse, error) {
    
}
func (s *server) UpdateUserDescription(ctx context.Context, in *pb.UpdateUserDescriptionRequest) (*pb.UpdateUserDescriptionResponse, error) {
    
}
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatal("Error", err)
    }

    s := grpc.NewServer()
    pb.RegisterUserDataMessageServiceServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}