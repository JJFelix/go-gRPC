package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/JJFelix/go-gRPC/proto"
)

const (
	port = ":8080"
)

func main(){
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Cucurella", "James", "Malo" },
	}

	// callSayHello(client) // Unary API
	// callSayHelloServerStream(client, names) // Server streaming
	// callSayHelloClientStream(client, names) // Client streaming
	callSayHelloBidirectionalStream(client, names) // Bidirectional streaming
}