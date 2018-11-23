package main

import(
	"fmt"
	"net"
	"google.golang.org/grpc"
	"github.com/venk3389/gRPC-Course/greet/greetpb"
)


func main(){

	cc, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil{
		fmt.Println(err.Error())
	}
	defer cc.Close()

	fmt.Println(greetpb.NewGreetServiceClient(cc))
}
