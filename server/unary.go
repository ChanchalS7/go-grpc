package main

import (
	"context"
	pb "github.com/ChanchalS7/go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error){
	return &pb.HelloResponse{
		Message:"Hello",
	},nil
}