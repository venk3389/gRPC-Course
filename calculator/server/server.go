package main

import(
	"context"
	"fmt"
	"io"
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
	N := in.GetNum()
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
			//time.Sleep(1000 * time.Millisecond)
			N = N / k      // divide N by k so that we have the rest of the number left.
		} else {
			k = k + 1
		}
	}
	return nil
}

func (*server) ComputeAverage(s sum.SumService_ComputeAverageServer) error{
	var avg float64 = 0.0
	var add int32 = 0
	var count int32 = 0
	for {
		count++
		num, err := s.Recv()
		if err == io.EOF{
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		add += num.GetNum()
		avg = float64(add)/float64(count)
	}

	return s.SendAndClose(&sum.ComputeAverageResponse{
		Average:avg,
	})
}

func (*server) ComputeMaximum(stream sum.SumService_ComputeMaximumServer) error{
	var largest int32
	for {
		data, err := stream.Recv()
		if err != nil || err == io.EOF{
			fmt.Println(err.Error())
			return err
		}
		if data.GetNum() > largest{
			largest = data.GetNum()
		}
		err = stream.Send(&sum.ComputeMaximumResponse{
			Num:largest,
		})
		if err != nil{
			fmt.Println(err.Error())
			return err
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
