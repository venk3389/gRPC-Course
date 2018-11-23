package main

import(
	"context"
	"fmt"
	"net"
	"github.com/venk3389/gRPC-Course/calculator/sumpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *sum.SumRequest) (*sum.SumResponse, error){
	fmt.Println(ctx)

	resp := &sum.SumResponse{
		Response : req.Req.GetFirstNumber() + req.Req.GetSecondNumber(),
	}
	return resp,nil
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
