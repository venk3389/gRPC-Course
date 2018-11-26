package main

import (
	"context"
	"fmt"
	"github.com/venk3389/gRPC-Course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/credentials"
	"net"
	"time"
)

type server struct{}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {

	p, ok := peer.FromContext(ctx)
	if ok {
		fmt.Println("Server received request:", req, " from ", p.Addr.String())
	}
	time.Sleep(5 * time.Second)
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	}

	name := req.Greeting.GetFirstName()
	if len(name) >= 10 || len(name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			"Length of `Name` cannot be more than 10 characters or empty")

	}
	return &greetpb.GreetWithDeadlineResponse{
		Response: "Hello " + name + "!",
	}, nil

}

func main() {
	// adding config for tls
	opts := []grpc.ServerOption{}
	tls := true
	fmt.Println("Running server in tls mode:",tls)
	if tls {
		certFile := "../tls/server.crt"
		keyFile := "../tls/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			fmt.Println("Failed loading certificates: ", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		fmt.Println(err.Error())
	}
	s := grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		fmt.Println(err.Error())
	}
}
