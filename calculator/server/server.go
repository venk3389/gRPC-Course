package main

import(
	"context"
	"time"
	"fmt"
	"net"
	"github.com/venk3389/gRPC-Course/calculator/sumpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sum.SumRequest) (*sum.SumResponse, error){
	fmt.Println(ctx)

	resp := &sum.SumResponse{
		Response : req.Req.GetFirstNumber() + req.Req.GetSecondNumber(),
	}
	return resp,nil
}

func (*server) PrimeNumberDecomposition(in *sum.PrimeNumberRequest, stream sum.SumService_PrimeNumberDecompositionServer) error{

	var k int32 = 2
	N := in.Num
	for N > 1 {
		if N%k == 0 { // if k evenly divides into N
			fmt.Println(N,k)
			resp := &sum.PrimeNumberResponse{
				Num:k,
			}// send the factor to the client
			if err := stream.Send(resp); err != nil{
				fmt.Println(err.Error())
				return err
			}
			time.Sleep(1000 * time.Millisecond)
			N = N / k      // divide N by k so that we have the rest of the number left.
		} else {
			k = k + 1
		}
	}
	return nil
}


func main(){
	lis, err := net.Listen("tcp",":50051")
	if err != nil{
		fmt.Println(err.Error())
	}
	s := grpc.NewServer()
	sum.RegisterSumServiceServer(s, &server{})
	if err = s.Serve(lis);err != nil{
		fmt.Println(err.Error())
	}
}
