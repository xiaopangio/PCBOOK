// Package service  @Author xiaobaiio 2023/2/21 15:24:00
package service_test

import (
	"bufio"
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/xiaopangio/pcbook/pb"
	"github.com/xiaopangio/pcbook/sample"
	"github.com/xiaopangio/pcbook/serializer"
	"github.com/xiaopangio/pcbook/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net"
	"os"
	"path/filepath"
	"testing"
)

func TestClientLaptapCreate(t *testing.T) {
	t.Parallel()
	laptapStore := service.NewInMemoryLaptapStore()
	serverAddr := startTestLaptapServer(t, laptapStore, nil, nil)
	laptapClient := newTestLaptapClient(t, serverAddr)
	laptap := sample.NewLaptap()
	expectedId := laptap.Id
	req := &pb.CreateLaptapRequest{Laptap: laptap}
	res, err := laptapClient.CreateLaptap(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.Id, expectedId)
}

func requireSameLaptap(t *testing.T, laptap1 *pb.Laptap, laptap2 *pb.Laptap) {
	json1, err := serializer.ProtobufToJson(laptap1)
	require.NoError(t, err)
	json2, err := serializer.ProtobufToJson(laptap2)
	require.NoError(t, err)
	require.Equal(t, json1, json2)
}

func newTestLaptapClient(t *testing.T, addr string) pb.LaptapServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return pb.NewLaptapServiceClient(conn)
}

func startTestLaptapServer(t *testing.T, laptapStore service.LaptapStore, imageStore service.ImageStore, ratingStore service.RatingStore) string {
	laptapServer := service.NewLaptapServer(laptapStore, imageStore, ratingStore)
	grpcServer := grpc.NewServer()
	pb.RegisterLaptapServiceServer(grpcServer, laptapServer)
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	go grpcServer.Serve(listener)
	return listener.Addr().String()
}
func TestClientLaptapSearch(t *testing.T) {
	t.Parallel()
	filter := &pb.Filter{
		MaxPriceUsd: 2000,
		MinCouCores: 4,
		MinCpuHz:    2.2,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}
	store := service.NewInMemoryLaptapStore()
	expectedIDs := make(map[string]bool)
	for i := 0; i < 6; i++ {
		laptap := sample.NewLaptap()
		switch i {
		case 0:
			laptap.PriceUsd = 2500
		case 1:
			laptap.GetCpu().NumbersCores = 2
		case 2:
			laptap.GetCpu().MinGhz = 2.0
		case 3:
			laptap.Ram = &pb.Memory{Value: 4096, Unit: pb.Memory_MEGABYTE}
		case 4:
			laptap.PriceUsd = 1999
			laptap.GetCpu().NumbersCores = 8
			laptap.GetCpu().MinGhz = 2.5
			laptap.GetCpu().MaxGhz = 4.5
			laptap.Ram = &pb.Memory{Value: 16, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptap.Id] = true
		case 5:
			laptap.PriceUsd = 1800
			laptap.GetCpu().NumbersCores = 6
			laptap.GetCpu().MinGhz = 2.3
			laptap.GetCpu().MaxGhz = 4.0
			laptap.Ram = &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptap.Id] = true
		}
		err := store.Save(laptap)
		require.NoError(t, err)
	}

	serverAddr := startTestLaptapServer(t, store, nil, nil)
	laptapClient := newTestLaptapClient(t, serverAddr)
	req := &pb.SearchLaptapRequest{Filter: filter}
	stream, err := laptapClient.SearchLaptap(context.Background(), req)
	require.NoError(t, err)
	found := 0
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		require.Contains(t, expectedIDs, res.GetLaptap().GetId())
		found++
	}
	require.Equal(t, len(expectedIDs), found)
}
func TestClientUploadImage(t *testing.T) {
	testImageFolder := "../tmp"
	laptapStore := service.NewInMemoryLaptapStore()
	imageStore := service.NewDiskImageStore(testImageFolder)
	serverAddr := startTestLaptapServer(t, laptapStore, imageStore, nil)
	laptap := sample.NewLaptap()
	err := laptapStore.Save(laptap)
	require.NoError(t, err)
	laptapClient := newTestLaptapClient(t, serverAddr)
	imagePath := fmt.Sprintf("%s/laptap.png", testImageFolder)
	file, err := os.Open(imagePath)
	require.NoError(t, err)
	defer file.Close()
	stream, err := laptapClient.UploadImage(context.Background())
	require.NoError(t, err)
	imageType := filepath.Ext(imagePath)
	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptapId:  laptap.GetId(),
				ImageType: imageType,
			},
		},
	}
	err = stream.Send(req)
	require.NoError(t, err)
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	size := 0
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		size += n
		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}
		err = stream.Send(req)
		require.NoError(t, err)
	}
	res, err := stream.CloseAndRecv()
	require.NoError(t, err)
	require.NotZero(t, size)
	require.Equal(t, size, int(res.GetSize()))
	savedPath := fmt.Sprintf("%s/%s%s", testImageFolder, res.GetId(), imageType)
	require.FileExists(t, savedPath)
	require.NoError(t, os.Remove(savedPath))
}
func TestClientRateLaptap(t *testing.T) {
	laptapStore := service.NewInMemoryLaptapStore()
	ratingStore := service.NewInMemoryRatingScore()
	serverAddr := startTestLaptapServer(t, laptapStore, nil, ratingStore)
	laptap := sample.NewLaptap()
	err := laptapStore.Save(laptap)
	require.NoError(t, err)
	laptapClient := newTestLaptapClient(t, serverAddr)
	stream, err := laptapClient.RateLaptap(context.Background())
	require.NoError(t, err)
	scores := []float64{8, 7.5, 10}
	averageScores := []float64{8, 7.75, 8.5}
	n := len(scores)
	for i := 0; i < n; i++ {
		req := &pb.RateLaptapRequest{
			Id:    laptap.GetId(),
			Score: scores[i],
		}
		err := stream.Send(req)
		require.NoError(t, err)
	}
	err = stream.CloseSend()
	require.NoError(t, err)
	for idx := 0; ; idx++ {
		res, err := stream.Recv()
		if err == io.EOF {
			require.Equal(t, n, idx)
			return
		}
		require.NoError(t, err)
		require.Equal(t, laptap.Id, res.GetId())
		require.Equal(t, res.GetRateCount(), uint32(idx+1))
		require.Equal(t, res.GetAverageScore(), averageScores[idx])
	}
}
