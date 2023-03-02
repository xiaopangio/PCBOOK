// Package service  @Author xiaobaiio 2023/2/22 18:38:00
package service

import "sync"

type RatingStore interface {
	Add(laptapId string, score float64) (*Rating, error)
}
type Rating struct {
	count uint32
	sum   float64
}
type InMemoryRatingScore struct {
	mutex  sync.RWMutex
	rating map[string]*Rating
}

func NewInMemoryRatingScore() *InMemoryRatingScore {
	return &InMemoryRatingScore{
		rating: make(map[string]*Rating),
	}
}

func (store *InMemoryRatingScore) Add(laptapId string, score float64) (*Rating, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	rating := store.rating[laptapId]
	if rating == nil {
		rating = &Rating{
			count: 1,
			sum:   score,
		}
	} else {
		rating.count++
		rating.sum += score
	}
	store.rating[laptapId] = rating
	return rating, nil
}
