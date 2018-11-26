package main

import (
	"context"
	"fmt"
	"github.com/venk3389/gRPC-Course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	//"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func doGreetWithDeadline(cc *grpc.ClientConn, first_name, last_name string, timeout uint32) {
	var cancel context.CancelFunc
	ctx := context.Background()
	if timeout != 0 {
		fmt.Println(fmt.Sprintf("Creating client request with deadline of %d ms", timeout))
		ctx, cancel = context.WithTimeout(ctx, time.Duration(timeout)*time.Millisecond)
		defer cancel()
	}
	s := greetpb.NewGreetServiceClient(cc)
	response, err := s.GreetWithDeadline(ctx, &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: first_name,
			LastName:  last_name,
		},
	})
	if err != nil {
		errStatus, ok := status.FromError(err)
		// lets print the error code which is `INVALID_ARGUMENT`
		if ok {
			//if codes.InvalidArgument == errStatus.Code() {
			// do your stuff here
			fmt.Println("Message:", errStatus.Message(), "Code:", errStatus.Code())
			//}
			return
		}
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Response from server: ", response.GetResponse())

}
func main() {
	fmt.Println("Running Secure gRPC greet client.....")
	tls := true
	opts := grpc.WithInsecure()
	if tls {
		certFile := "../tls/ca.crt"
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			fmt.Println("Error loading the ca certificate file,", sslErr.Error())
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer cc.Close()
	//greet call that returns a deadline exceeded 
	doGreetWithDeadline(cc, "", "jdhdcb", 1000)
	//greet call that returns a valid response
	doGreetWithDeadline(cc, "Venkat", "Tester", 0)
	//greet call that returns a invalid argument response
	doGreetWithDeadline(cc, "Venkatsjdbksjdb", "Tester", 0)
}
