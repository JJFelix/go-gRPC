package main

import (
	"log"
	"net"
	"google.golang.org/grpc"

	pb "github.com/JJFelix/go-gRPC/proto"
)

const (
	port = ":8080"
)

type helloServer struct{
	pb.GreetServiceServer
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("Failed to start the server: %v ", err)
	}

	grpcServer := grpc.NewServer()
	
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Failed to start: %v", err)
	}

}