package main

import (
	"flag"
	"log"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	pb "github.com/gust1n/randomNumberService/randomnumber"
)

var address = flag.String("addr", "localhost:9000", "Server address")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewRandomNumberClient(conn)

	for {
		r, err := c.GetInt(context.Background(), &pb.GetIntRequest{Max: int32(1000)})
		if err != nil {
			log.Fatalf("could not get random int: %v", err)
		}
		log.Printf("Random int: %d", r.RandomInt)
		time.Sleep(time.Millisecond * 500)
	}
}
