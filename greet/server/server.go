package main

import(
	"context"
	"fmt"
	"net"
	"github.com/venk3389/gRPC-Course/greet/greetpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetWithDeadline(ctx context.Context,req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error){
	name := req.Greeting.GetFirstName()
	if len(name) >= 10 || len(name) == 0{
		return nil, status.Errorf(codes.InvalidArgument,
			"Length of `Name` cannot be more than 10 characters or empty")

	}
	response := "Hello "+name+"!"
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
