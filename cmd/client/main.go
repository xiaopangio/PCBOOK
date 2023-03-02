// Package client  @Author xiaobaiio 2023/2/21 16:04:00
package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/xiaopangio/pcbook/client"
	"github.com/xiaopangio/pcbook/pb"
	"github.com/xiaopangio/pcbook/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strings"
	"time"
)

func testCreateLaptap(laptapClient *client.LaptapClient) {
	laptapClient.CreateLaptap(sample.NewLaptap())
}
func testSearchLaptap(laptapClient *client.LaptapClient) {
	for i := 0; i < 10; i++ {
		laptapClient.CreateLaptap(sample.NewLaptap())
	}
	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCouCores: 4,
		MinCpuHz:    2.2,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}
	laptapClient.SearchLaptap(filter)
}
func testUploadImage(laptapClient *client.LaptapClient) {
	laptap := sample.NewLaptap()
	laptapClient.CreateLaptap(laptap)
	laptapClient.UploadImage(laptap.GetId(), "tmp/laptap.png")
}
func testRateLaptap(laptapClient *client.LaptapClient) {
	n := 3
	laptapIDs := make([]string, n)
	for i := 0; i < n; i++ {
		laptap := sample.NewLaptap()
		laptapIDs[i] = laptap.Id
		laptapClient.CreateLaptap(laptap)
	}
	scores := make([]float64, n)
	for {
		fmt.Print("rate laptap (y/n)?")
		var answer string
		fmt.Scan(&answer)
		if strings.ToLower(answer) != "y" {
			break
		}
		for i := 0; i < n; i++ {
			scores[i] = sample.RandomLaptapScore()
		}
		err := laptapClient.RateLaptap(laptapIDs, scores)
		if err != nil {
			log.Fatal(err)
		}
	}

}

const (
	username        = "admin1"
	password        = "secret"
	refreshDuration = 30 * time.Second
)

func authMethods() map[string]bool {
	const laptapServicePath = "/techschool.pcbook.LaptapService/"
	return map[string]bool{
		laptapServicePath + "CreateLaptap": true,
		laptapServicePath + "UploadImage":  true,
		laptapServicePath + "RateLaptap":   true,
	}
}
func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemServeCA, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServeCA) {
		return nil, fmt.Errorf("failed to add CA's certificate")
	}
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}
	return credentials.NewTLS(config), err
}
func main() {
	serverAddress := flag.String("address", "", "the server address")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()
	log.Printf("dial server %s, TLS = %t", *serverAddress, *enableTLS)
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	if *enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		transportOption = grpc.WithTransportCredentials(tlsCredentials)
	}

	cc1, err := grpc.Dial(*serverAddress, transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	authClient := client.NewAuthClient(cc1, username, password)
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatalf("cannot create auth interceptor: %v", err)
	}
	cc2, err := grpc.Dial(
		*serverAddress,
		transportOption,
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	laptapClient := client.NewLaptapClient(cc2)
	testRateLaptap(laptapClient)
}
