package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"path/filepath"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"

	pb "github.com/gust1n/randomNumberService/randomnumber"
)

type randomNumberServer struct{}

func (s *randomNumberServer) GetInt(ctx context.Context, req *pb.GetIntRequest) (*pb.GetIntResponse, error) {
	randInt := int32(rand.Intn(int(req.Max)))
	if randInt > 800 {
		return nil, grpc.Errorf(codes.OutOfRange, "random cannot be too big")
	}
	return &pb.GetIntResponse{randInt}, nil

}

func withConfigDir(path string) string {
	return filepath.Join(os.Getenv("HOME"), ".randomNumber", "server", path)
}

func main() {
	var (
		caCert     = flag.String("ca-cert", withConfigDir("ca.pem"), "Trusted CA certificate.")
		listenAddr = flag.String("listen-addr", "0.0.0.0:9000", "HTTP listen address.")
		tlsCert    = flag.String("tls-cert", withConfigDir("cert.pem"), "TLS server certificate.")
		tlsKey     = flag.String("tls-key", withConfigDir("key.pem"), "TLS server key.")
	)
	flag.Parse()

	log.Println("Random number service starting...")

	cert, err := tls.LoadX509KeyPair(*tlsCert, *tlsKey)
	if err != nil {
		log.Fatal(err)
	}

	rawCaCert, err := ioutil.ReadFile(*caCert)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(rawCaCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	gs := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterRandomNumberServer(gs, &randomNumberServer{})

	lis, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Random service started successfully")
	log.Fatal(gs.Serve(lis))
}
