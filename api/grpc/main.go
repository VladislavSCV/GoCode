package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/VladislavSCV/GoCode/api/grpc/gen/pb-go/com.user_data"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserDataMessageServiceClient(conn)

	r, err := c.UpdateUserPassword(context.Background(), &pb.UpdateUserPasswordRequest{UserId: 2, NewPassword: "1234567890"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Result: %s", r.String())
}
