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

	doComputeMaximum(cc)
}

func doComputeMaximum(cc *grpc.ClientConn){
	s := sum.NewSumServiceClient(cc)
	waitCh :=  make(chan struct{})//channel to wait on data from the server
	stream, err := s.ComputeMaximum(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	// goroutine to send data to server
	go func(){
		nums := []int{5,2,1,3,6,9}
		for _,v := range nums{
			if err := stream.Send(&sum.ComputeMaximumRequest{
				Num: int32(v),
			});err != nil || err == io.EOF{
				fmt.Println(err.Error())
				break
			}
		}
		stream.CloseSend()
	}()
	//goroutine to receive data from server
	go func(){
		for {
			resp,err := stream.Recv()
			if err == io.EOF{
				break
			}
			if err != nil{
				fmt.Println(err.Error())
				break
			}
			fmt.Println("The current maximum: ",resp.GetNum())
		}
		close(waitCh)	
	}()
	<-waitCh // exit on receiving all the data from server
}

func doComputeAverage(cc *grpc.ClientConn) {
	s := sum.NewSumServiceClient(cc)
	vals := []int32{1, 2, 3, 4}
	stream, err := s.ComputeAverage(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range vals {
		s := sum.ComputeAverageRequest{
			Num: v,
		}
		if err = stream.Send(&s); err != nil {
			fmt.Println(err.Error())
		}
	}
	average,err := stream.CloseAndRecv()
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println("Average returned is:",average.GetAverage())
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
