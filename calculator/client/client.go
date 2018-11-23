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
	l := &sum.SumRequest{
		Req : &sum.Request{
			FirstNumber : 1241,
			SecondNumber : 265,
		      },
	}
	response,err := s.Sum(context.Background(),l)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(response.Response)
}
