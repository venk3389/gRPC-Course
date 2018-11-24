package main

import (
	"context"
	"fmt"
	"github.com/venk3389/gRPC-Course/calculator/sumpb"
	"google.golang.org/grpc"
	"io"
)

func main() {

	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer cc.Close()

	doPrimeDecomposition(cc)
}

func doPrimeDecomposition(cc *grpc.ClientConn) {
	s := sum.NewSumServiceClient(cc)
	l := &sum.PrimeNumberRequest{
		Num: 210345670,
	}
	response, err := s.PrimeNumberDecomposition(context.Background(), l)
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		out, err := response.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(out.GetNum())
	}

}
func doSum(cc *grpc.ClientConn) {

	s := sum.NewSumServiceClient(cc)
	l := &sum.SumRequest{
		Req: &sum.Request{
			FirstNumber:  1241,
			SecondNumber: 265,
		},
	}
	response, err := s.Sum(context.Background(), l)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(response.Response)
}
