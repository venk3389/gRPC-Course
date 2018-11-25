package main

import (
	"fmt"
	"github.com/venk3389/gRPC-Course/greet/greetpb"
	"google.golang.org/grpc"
	"context"
)

func main() {

	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer cc.Close()

	s := greetpb.NewGreetServiceClient(cc)
	response, err := s.GreetWithDeadline(context.Background(), &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "venkat",
			LastName:  "lester",
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Response from server: ", response.GetResponse())
}
