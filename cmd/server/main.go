// Package server  @Author xiaobaiio 2023/2/21 16:04:00
package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/xiaopangio/pcbook/pb"
	"github.com/xiaopangio/pcbook/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"time"
)

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

func seedUsers(userStore service.UserStore) error {
	err := createUser(userStore, "admin1", "secret", "admin")
	if err != nil {
		return err
	}
	return createUser(userStore, "user1", "secret", "user")
}
func createUser(userStore service.UserStore, username, password, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}
func accessibleRoles() map[string][]string {
	const laptapServicePath = "/techschool.pcbook.LaptapService/"
	return map[string][]string{
		laptapServicePath + "CreateLaptap": {"admin"},
		laptapServicePath + "UploadImage":  {"admin"},
		laptapServicePath + "RateLaptap":   {"admin", "user"},
	}
}
func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemClientCA, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequestClientCert,
		ClientCAs:    certPool,
	}
	return credentials.NewTLS(config), err
}
func main() {
	//flag
	port := flag.Int("port", 0, "the server port")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()
	log.Printf("start server on port %d, TLS = %t", *port, *enableTLS)
	//laptop
	laptapStore := service.NewInMemoryLaptapStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingScore()
	laptapServer := service.NewLaptapServer(laptapStore, imageStore, ratingStore)
	//auth
	userStore := service.NewInMemoryUserStore()
	err := seedUsers(userStore)
	if err != nil {
		log.Fatal("cannot seed users")
	}
	jwtManager := service.NewJWTManager(secretKey, tokenDuration)
	authServer := service.NewAuthServer(userStore, jwtManager)
	interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())
	//server
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	if *enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}
	grpcServer := grpc.NewServer(serverOptions...)
	//register
	pb.RegisterLaptapServiceServer(grpcServer, laptapServer)
	pb.RegisterAuthServiceServer(grpcServer, authServer)
	//reflection
	reflection.Register(grpcServer)
	//listen
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
