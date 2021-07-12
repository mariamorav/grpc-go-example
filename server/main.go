package main

import (
	"context"
	"fmt"
	pb "github.com/mariamorav/grpc-go-example/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedWishListServiceServer
}

func (s *server) create(ctx context.Context, req *pb.CreateWishListRequest) (*pb.CreateWishListResponse, error) {
	fmt.Println("creating the wishing list..." + req.WishList.Name)
	return &pb.CreateWishListResponse{
		WishListId: req.WishList.Id,
	}, nil
}

func (s *server) Add(context.Context, *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	return nil, nil
}

func (s *server) List(context.Context, *pb.ListWishListRequest) (*pb.ListWishListResponse, error) {
	return nil, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tpc connection: " + err.Error())
	}

	serve := grpc.NewServer()
	pb.RegisterWishListServiceServer(serve, &server{})
	if err = serve.Serve(listener); err != nil {
		panic("cannot initialize the server: " + err.Error())
	}

	fmt.Println("Server running on port 50051")
}
