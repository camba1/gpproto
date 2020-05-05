package main

import (
	"context"
	"google.golang.org/grpc"
	pb "gpProto/product"
	"log"
	"net"
)

const port = ":50051"

type prodServer struct {
	pb.UnimplementedProductSrvServer
}

func (p *prodServer) GetProduct(ctx context.Context, saerchParms *pb.SearchParams) (*pb.Product, error) {
	_ = ctx
	prod := &pb.Product{
		Id:          saerchParms.Id,
		Name:        saerchParms.Name,
		Description: "test me",
	}
	return prod, nil
}

func (p *prodServer) CreateProduct(ctx context.Context, product *pb.Product) (*pb.Product, error) {
	_ = ctx
	return product, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Cannot start listening on port %s. Error: %v", port, err)
	}
	s := grpc.NewServer()
	pb.RegisterProductSrvServer(s, &prodServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Unable to start product server. Error: %v", err)
	}
}
