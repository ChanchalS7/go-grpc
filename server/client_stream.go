package main

import(
	"io"
	"log"
	pb "github.com/ChanchalS7/go-grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error{
	var message []string

	for{
		req, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&pb.MessageList{Messages: message})
		}
		if err!=nil{
			return err
		}
		log.Printf("Got request with name:%v",req.Name)
		message=append(message, "Hello",req.Name)
	}
}