// Package client  @Author xiaobaiio 2023/2/23 15:52:00
package client

import (
	"bufio"
	"context"
	"fmt"
	"github.com/xiaopangio/pcbook/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type LaptapClient struct {
	service pb.LaptapServiceClient
}

func NewLaptapClient(conn *grpc.ClientConn) *LaptapClient {
	return &LaptapClient{service: pb.NewLaptapServiceClient(conn)}
}

func (laptapClient *LaptapClient) CreateLaptap(laptap *pb.Laptap) {
	req := &pb.CreateLaptapRequest{Laptap: laptap}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := laptapClient.service.CreateLaptap(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("laptap already exists")
		} else {
			log.Fatal("cannot create a laptap: ", err)
		}
		return
	}
	log.Printf("created laptap with id: %s", res.Id)
}
func (laptapClient *LaptapClient) SearchLaptap(filter *pb.Filter) {
	log.Print("search filter: ", filter)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.SearchLaptapRequest{Filter: filter}
	stream, err := laptapClient.service.SearchLaptap(ctx, req)
	if err != nil {
		log.Fatal("cannot search laptap: ", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}
		laptap := res.GetLaptap()
		log.Print("- found: ", laptap.GetId())
		log.Print("  + brand: ", laptap.GetBrand())
		log.Print("  + name: ", laptap.GetName())
		log.Print("  + cpu cores: ", laptap.GetCpu().GetNumbersCores())
		log.Print("  + cpu min ghz: ", laptap.GetCpu().GetMinGhz())
		log.Print("  + ram: ", laptap.GetRam().GetValue(), laptap.GetRam().GetUnit())
		log.Print("  + price: ", laptap.GetPriceUsd(), "usd")
	}
}

func (laptapClient *LaptapClient) UploadImage(laptapId string, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := laptapClient.service.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot send uploadImage request: ", err)
	}
	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptapId:  laptapId,
				ImageType: filepath.Ext(imagePath),
			},
		},
	}
	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info: ", err, stream.RecvMsg(nil))
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}
		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk data to server: ", err, stream.RecvMsg(nil))
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}
	log.Printf("image uploaded with id: %s, size: %d", res.GetId(), res.GetSize())
}
func (laptapClient *LaptapClient) RateLaptap(laptapIDs []string, laptapScores []float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := laptapClient.service.RateLaptap(ctx)
	if err != nil {
		return fmt.Errorf("cannot send rateLaptap request: %v", err)
	}
	//wait response
	waitResponse := make(chan error)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Print("no more response")
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("cannot receive stream response: %v", err)
				return
			}
			log.Print("received response: ", res)
		}
	}()
	for i, laptapID := range laptapIDs {
		req := &pb.RateLaptapRequest{
			Id:    laptapID,
			Score: laptapScores[i],
		}
		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send stream request: %v - %v", err, stream.RecvMsg(nil))
		}
		log.Print("send request: ", req)
	}
	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("cannot close send: %v", err)
	}
	err = <-waitResponse
	return err
}
