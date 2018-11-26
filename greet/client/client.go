package main

import (
	"context"
	"fmt"
	"github.com/venk3389/gRPC-Course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
func doGreetWithDeadline(cc *grpc.ClientConn, first_name,last_name string){

s := greetpb.NewGreetServiceClient(cc)
        response, err := s.GreetWithDeadline(context.Background(), &greetpb.GreetWithDeadlineRequest{
                Greeting: &greetpb.Greeting{
                        FirstName: first_name,
                        LastName:  last_name,
                },
        })
        if err != nil {
                errStatus, ok := status.FromError(err)
                // lets print the error code which is `INVALID_ARGUMENT`
                if ok {
                        if codes.InvalidArgument == errStatus.Code() {
                                // do your stuff here
                                fmt.Println("Message:", errStatus.Message(), "Code:", errStatus.Code())
                        }
                        return
                }
                fmt.Println(err.Error())
                return
        }
        fmt.Println("Response from server: ", response.GetResponse())



}
func main() {
	fmt.Println("Running gRPC greet client.....")
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer cc.Close()
	//greet call that returns an error message
	doGreetWithDeadline(cc,"","jdhdcb")
	//greet call that returns a valid response
	doGreetWithDeadline(cc,"Venkat","Tester")
}
