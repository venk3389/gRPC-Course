package main

import(
	"fmt"
	"google.golang.org/grpc"
	"github.com/venk3389/gRPC-Course/calculator/sumpb"
	"context"
)


func main(){

	cc, err := grpc.Dial("0.0.0.0:50051",grpc.WithInsecure())
	if err != nil{
		fmt.Println(err.Error())
	}
	defer cc.Close()

	s := sum.NewSumServiceClient(cc)
	l := &sum.Request{
		FirstNumber : 24,
		SecondNumber : 26,
	}
	response := s.Sum(context.Background(),l)
	fmt.Println(response.Response)
}
