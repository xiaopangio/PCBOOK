// Package service  @Author xiaobaiio 2023/2/21 13:03:00
package service

import (
	"bytes"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/xiaopangio/pcbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

const maxImageSize = 1 << 20

type LaptapServer struct {
	laptapStore LaptapStore
	imageStore  ImageStore
	ratingStore RatingStore
	pb.UnimplementedLaptapServiceServer
}

func NewLaptapServer(laptapStore LaptapStore, imageStore ImageStore, ratingStore RatingStore) *LaptapServer {
	return &LaptapServer{laptapStore, imageStore, ratingStore, pb.UnimplementedLaptapServiceServer{}}
}

func (server *LaptapServer) CreateLaptap(
	ctx context.Context,
	request *pb.CreateLaptapRequest,
) (*pb.CreateLaptapResponse, error) {
	laptap := request.GetLaptap()
	log.Printf("receive a create-laptap request with id: %s", laptap.Id)
	if len(laptap.Id) > 0 {
		//check if it's a valid uuid
		_, err := uuid.Parse(laptap.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptap ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptap ID: %v", err)
		}
		laptap.Id = id.String()
	}

	//some heavy processing
	//time.Sleep(6 * time.Second)
	//check context
	if err := contextError(ctx); err != nil {
		return nil, err
	}
	//save laptap to laptapStore
	err := server.laptapStore.Save(laptap)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptap to laptapStore: %v", err)
	}
	log.Printf("saved laptap with id: %v", laptap.Id)
	res := &pb.CreateLaptapResponse{
		Id: laptap.Id,
	}
	return res, nil

}
func (server *LaptapServer) SearchLaptap(
	req *pb.SearchLaptapRequest,
	stream pb.LaptapService_SearchLaptapServer,
) error {
	filter := req.GetFilter()
	log.Printf("receive a search-laptap request with filter: %v", filter)
	err := server.laptapStore.Search(
		stream.Context(),
		filter,
		func(laptap *pb.Laptap) error {
			res := &pb.SearchLaptapResponse{Laptap: laptap}
			err := stream.Send(res)
			if err != nil {
				return err
			}
			log.Printf("send laptap with id: %s", laptap.GetId())
			return nil
		},
	)
	if err != nil {
		return status.Errorf(codes.Internal, "unexpected err: %v", err)
	}
	return nil
}
func (server *LaptapServer) UploadImage(stream pb.LaptapService_UploadImageServer) error {
	recv, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}
	laptapId := recv.GetInfo().GetLaptapId()
	imageType := recv.GetInfo().GetImageType()
	log.Printf("receive an upload-image request for laptap %s with image type %s", laptapId, imageType)
	laptap, err := server.laptapStore.Find(laptapId)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot find laptap: %s", err))
	}
	if laptap == nil {
		return logError(status.Errorf(codes.InvalidArgument, "laptap %s doesn't exist", laptapId))
	}
	imageData := bytes.Buffer{}
	imageSize := 0
	for {
		//check context
		if err = contextError(stream.Context()); err != nil {
			return err
		}
		log.Print("waiting to receive more data")
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %s", err))
		}
		chunk := req.GetChunkData()
		size := len(chunk)
		log.Printf("receivce a chunk with size: %d", size)
		imageSize += size
		if imageSize > maxImageSize {
			return logError(status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize))
		}
		_, err = imageData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %s", err))
		}
	}
	imageId, err := server.imageStore.Save(laptapId, imageType, imageData)
	res := &pb.UploadImageResponse{
		Id:   imageId,
		Size: uint32(imageSize),
	}
	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response: %s", err))
	}
	log.Printf("save image with id: %s, size: %d", imageId, imageSize)
	return nil
}
func (server *LaptapServer) RateLaptap(stream pb.LaptapService_RateLaptapServer) error {
	for {
		//check context
		err := contextError(stream.Context())
		if err != nil {
			return err
		}
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		laptapId := req.GetId()
		score := req.GetScore()
		log.Printf("received a rate-laptap request: id = %s, score = %.2f", laptapId, score)
		found, err := server.laptapStore.Find(laptapId)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot find laptap: %v", err))
		}
		if found == nil {
			return logError(status.Errorf(codes.NotFound, "laptap %s is not exist", laptapId))
		}
		rating, err := server.ratingStore.Add(laptapId, score)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot add rating to the store: %s", err))
		}
		res := &pb.RateLaptapResponse{
			Id:           laptapId,
			RateCount:    rating.count,
			AverageScore: rating.sum / float64(rating.count),
		}
		err = stream.Send(res)
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
	}
	return nil
}
func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Errorf(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Errorf(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}
