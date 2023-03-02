// Package service  @Author xiaobaiio 2023/2/21 13:13:00
package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/xiaopangio/pcbook/pb"
	"log"
	"sync"
)

var ErrAlreadyExists = errors.New("record already exists")

type LaptapStore interface {
	Save(laptap *pb.Laptap) error
	Find(id string) (*pb.Laptap, error)
	Search(ctx context.Context, filter *pb.Filter, found func(laptap *pb.Laptap) error) error
}
type InMemoryLaptapStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptap
}

func NewInMemoryLaptapStore() *InMemoryLaptapStore {
	return &InMemoryLaptapStore{
		data: make(map[string]*pb.Laptap),
	}
}

func (store *InMemoryLaptapStore) Save(laptap *pb.Laptap) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if store.data[laptap.Id] != nil {
		return ErrAlreadyExists
	}
	other, err := deepCopy(laptap)
	if err != nil {
		return err
	}
	store.data[other.Id] = other
	return nil
}
func (store *InMemoryLaptapStore) Find(id string) (*pb.Laptap, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	laptap := store.data[id]
	if laptap == nil {
		return nil, nil
	}
	return deepCopy(laptap)
}
func (store *InMemoryLaptapStore) Search(ctx context.Context, filter *pb.Filter, found func(laptap *pb.Laptap) error) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	for _, laptap := range store.data {
		//time.Sleep(1 * time.Second)
		log.Print("check laptap: ", laptap.GetId())
		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Print("context is canceled")
			return errors.New("context is canceled")
		}
		if isQualified(filter, laptap) {
			other, err := deepCopy(laptap)
			if err != nil {
				return err
			}
			err = found(other)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isQualified(filter *pb.Filter, laptap *pb.Laptap) bool {
	if laptap.GetPriceUsd() > filter.GetMaxPriceUsd() {
		return false
	}
	if laptap.GetCpu().GetNumbersCores() < filter.GetMinCouCores() {
		return false
	}
	if laptap.GetCpu().GetMinGhz() < filter.GetMinCpuHz() {
		return false
	}
	if toBit(laptap.GetRam()) < toBit(filter.GetMinRam()) {
		return false
	}
	return true
}

func toBit(ram *pb.Memory) uint64 {
	value := ram.GetValue()
	switch ram.GetUnit() {
	case pb.Memory_BIT:
		return value
	case pb.Memory_BYTE:
		return value << 3
	case pb.Memory_KILOBYTE:
		return value << 13
	case pb.Memory_MEGABYTE:
		return value << 23
	case pb.Memory_GIGABYTE:
		return value << 33
	case pb.Memory_TERABYTE:
		return value << 43
	default:
		return 0
	}
}
func deepCopy(laptap *pb.Laptap) (*pb.Laptap, error) {
	other := &pb.Laptap{}
	err := copier.Copy(other, laptap)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptap data: %w", err)
	}
	return other, nil
}
