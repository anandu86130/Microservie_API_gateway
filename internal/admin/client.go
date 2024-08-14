package admin

import (
	"log"

	adminpb "github.com/Microservice_API_gateway/internal/admin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (adminpb.AdminServiceClient,error) {
	grpc, err := grpc.NewClient(":8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Printf("error when dialing to gRPC client: 8084")
		return nil,err
	}
	log.Printf("successfully connected to admin server at port: 8084")
	return adminpb.NewAdminServiceClient(grpc),nil
}