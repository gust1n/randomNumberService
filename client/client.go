package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/gust1n/randomNumberService/randomnumber"
)

func withConfigDir(path string) string {
	return filepath.Join(os.Getenv("HOME"), ".randomNumber", "client", path)
}

func main() {
	var (
		caCert     = flag.String("ca-cert", withConfigDir("ca.pem"), "Trusted CA certificate.")
		serverAddr = flag.String("server-addr", "127.0.0.1:9000", "Hello service address.")
		tlsCert    = flag.String("tls-cert", withConfigDir("cert.pem"), "TLS server certificate.")
		tlsKey     = flag.String("tls-key", withConfigDir("key.pem"), "TLS server key.")
	)
	flag.Parse()

	cert, err := tls.LoadX509KeyPair(*tlsCert, *tlsKey)
	if err != nil {
		log.Fatal(err)
	}

	rawCACert, err := ioutil.ReadFile(*caCert)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(rawCACert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	})

	conn, err := grpc.Dial(*serverAddr,
		grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
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
