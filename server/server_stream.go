package main

import (
	"time"
	"log"

	pb "github.com/ChanchalS7/go-grpc/proto"
)



func (s *helloServer) callSayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error{
log.Printf("Got request with name: %v",req.Names)

for _, name := range req.Names{
	res :=&pb.HelloResponse{
		Message: "Hello"+name,
	}
	if err := stream.Send(res); err!=nil{
		return err
	}
	time.Sleep(2 *time.Second)

}
return nil

}