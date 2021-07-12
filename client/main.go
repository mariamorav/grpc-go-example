package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-go-example/proto"
	"math/rand"
	"strconv"
	"time"
)

func generateID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewWishListServiceClient(conn)

	res, err := serviceClient.Create(context.Background(), &pb.CreateWishListRequest{
		WishList: &pb.WishList{
			Id:   generateID(),
			Name: "Standing desk",
		},
	})

	if err != nil {
		panic("wishlist is not created: " + err.Error())
	}

	fmt.Println(res.WishListId)
}
