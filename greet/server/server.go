package main

import(
	"fmt"
	"net"
	"github.com/venk3389/gRPC-Course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main(){
	lis, err := net.Listen("tcp",":50051")
	if err != nil{
		fmt.Println(err.Error())
	}	
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err = s.Serve(lis);err != nil{
		fmt.Println(err.Error())
	}
}
