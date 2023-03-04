// Package server  @Author xiaobaiio 2023/2/21 16:04:00
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/xiaopangio/pcbook/orm"
	"github.com/xiaopangio/pcbook/pb"
	"github.com/xiaopangio/pcbook/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	secretKey        = "secret"
	tokenDuration    = 15 * time.Minute
	serverCertFile   = "cert/server-cert.pem"
	serverKeyFile    = "cert/server-key.pem"
	clientCACertFile = "cert/ca-cert.pem"
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
	pemClientCA, err := os.ReadFile(clientCACertFile)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	serverCert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
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
func runGRPCServer(
	laptapServer pb.LaptapServiceServer,
	authServer pb.AuthServiceServer,
	jwtManager *service.JWTManager,
	enableTLS bool,
	listener net.Listener,
) error {
	interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())
	//server
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	if enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			return fmt.Errorf("cannot load TLS credentials: %w", err)
		}
		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}
	grpcServer := grpc.NewServer(serverOptions...)
	//register
	pb.RegisterLaptapServiceServer(grpcServer, laptapServer)
	pb.RegisterAuthServiceServer(grpcServer, authServer)
	//reflection
	reflection.Register(grpcServer)
	err := grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("cannot start server: %w", err)
	}
	log.Printf("start grpc server at %s, TLS = %t", listener.Addr(), enableTLS)
	return nil
}
func runRESTServer(
	enableTLS bool,
	listener net.Listener,
	endPoint string,
) error {
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterLaptapServiceHandlerFromEndpoint(ctx, mux, endPoint, opts)
	if err != nil {
		return fmt.Errorf("cannot register laptapServiceHandler from endpoint: %w", err)
	}
	err = pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, endPoint, opts)
	if err != nil {
		return fmt.Errorf("cannot register authServiceHandler from endpoint: %w", err)
	}
	log.Printf("start rest server at %s, TLS = %t", listener.Addr(), enableTLS)
	if enableTLS {
		err := http.ServeTLS(listener, mux, serverCertFile, serverKeyFile)
		return fmt.Errorf("cannot start https server: %w", err)
	} else {
		err := http.Serve(listener, mux)
		return fmt.Errorf("cannot start http server: %w", err)
	}
}
func main() {
	//flag
	port := flag.Int("port", 0, "the server port")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	grpcType := flag.String("type", "grpc", "type (grpc/rest)")
	endpoint := flag.String("endpoint", "", "endpoint")
	flag.Parse()
	//laptop
	laptapStore := service.NewInMemoryLaptapStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingScore()
	laptapServer := service.NewLaptapServer(laptapStore, imageStore, ratingStore)
	//auth
	//userStore := service.NewInMemoryUserStore()
	userStore := service.NewDBUserStore(orm.DB)

	jwtManager := service.NewJWTManager(secretKey, tokenDuration)
	authServer := service.NewAuthServer(userStore, jwtManager)
	//listen
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	if *grpcType == "grpc" {
		err = runGRPCServer(laptapServer, authServer, jwtManager, *enableTLS, listener)
		if err != nil {
			log.Fatal(err)
		}
		err := seedUsers(userStore)
		if err != nil {
			log.Fatal("cannot seed users: ", err)
		}
	} else {
		err := runRESTServer(*enableTLS, listener, *endpoint)
		if err != nil {
			log.Fatal(err)
		}
	}
}
