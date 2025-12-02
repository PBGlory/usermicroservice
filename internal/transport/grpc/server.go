package grpc

import (
	"log"
	"net"

	userpb "github.com/PBGlory/project-protos/proto/user"
	"github.com/PBGlory/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc *user.Service) error {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcSrv := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))

	log.Println("gRPC server started on :50051")

	return grpcSrv.Serve(l)
}
