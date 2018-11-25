package main

import(
	"context"
	"fmt"
	"net"
	"github.com/venk3389/gRPC-Course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetWithDeadline(ctx context.Context,req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error){
	response := "Hello "+req.Greeting.GetFirstName()+"!"
	return &greetpb.GreetWithDeadlineResponse{
		Response:response,
	},nil

}


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
