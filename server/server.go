package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	pb "github.com/gust1n/randomNumberService/randomnumber"
)

var port = flag.Int("port", 9000, "Port to serve on")

type randomNumberServer struct{}

func (s *randomNumberServer) GetInt(ctx context.Context, req *pb.GetIntRequest) (*pb.GetIntResponse, error) {
	randInt := int32(rand.Intn(int(req.Max)))
	if randInt > 800 {
		return nil, grpc.Errorf(codes.OutOfRange, "random cannot be too big")
	}
	return &pb.GetIntResponse{randInt}, nil

}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRandomNumberServer(grpcServer, &randomNumberServer{})
	grpcServer.Serve(lis)
}
