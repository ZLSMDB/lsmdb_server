package service

import (
	"context"
	"fmt"
	"io"

	pb "github.com/ZLSMDB/lsmdb_server/api/javaDemo/v1"
)

type JavaDemoServiceService struct {
	pb.UnimplementedJavaDemoServiceServer
}

func NewJavaDemoServiceService() *JavaDemoServiceService {
	return &JavaDemoServiceService{}
}

func (s *JavaDemoServiceService) HelloWorld(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{}, nil
}
func (s *JavaDemoServiceService) HelloWorldClientStream(conn pb.JavaDemoService_HelloWorldClientStreamServer) error {
	for {
		req, err := conn.Recv()
		if req != nil {
			fmt.Println(req.Request)
		}

		if err == io.EOF {
			return conn.SendAndClose(&pb.HelloResponse{})
		}
		if err != nil {
			return err
		}
	}
}
func (s *JavaDemoServiceService) HelloWorldClientAndServerStream(conn pb.JavaDemoService_HelloWorldClientAndServerStreamServer) error {
	for {
		req, err := conn.Recv()
		fmt.Println(req.Request)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		err = conn.Send(&pb.HelloResponse{})
		if err != nil {
			return err
		}
	}
}
