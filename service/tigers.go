package service

// go:generate mockgen -source=./service/tigers.go -destination=./mocks/mock_tigers.go -package=mocks github.com/pranoyk/tiger-sightings/service Tiger

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	customerr "github.com/pranoyk/tiger-sightings/custom-err"
	"github.com/pranoyk/tiger-sightings/model"
	"github.com/pranoyk/tiger-sightings/repository"
)

type tiger struct {
	repo repository.TigersRepository
}

type Tiger interface {
	CreateTiger(context.Context, *model.CreateTigerRequest, string) *customerr.APIError
	CreateSighting(context.Context, *model.CreateTigerSightingRequest, string) *customerr.APIError
}

func NewTiger(repo repository.TigersRepository) Tiger {
	return &tiger{
		repo: repo,
	}
}

func (t *tiger) CreateTiger(ctx context.Context, createTigerReq *model.CreateTigerRequest, email string) *customerr.APIError {
	parsedDob, err := time.Parse("2006-01-02", createTigerReq.Dob)
	if err != nil {
		return customerr.GetInvalidTimeError()
	}
	tiger := &model.Tiger{
		Name: createTigerReq.Name,
		Dob:  parsedDob,
	}
	parsedTime, err := time.Parse("2006-01-02T15:04:05Z", createTigerReq.LastSeen)
	if err != nil {
		return customerr.GetInvalidTimeError()
	}
	lat, err := strconv.ParseFloat(createTigerReq.Lat, 64)
	if err != nil {
		return &customerr.APIError{
			StatusCode: 400,
			Message:    "Invalid lat",
		}
	}
	lng, err := strconv.ParseFloat(createTigerReq.Lng, 64)
	if err != nil {
		return &customerr.APIError{
			StatusCode: 400,
			Message:    "Invalid lng",
		}
	}
	tigerSightings := &model.TigerSightings{
		LastSeen: parsedTime,
		Lat:      lat,
		Lng:      lng,
	}
	err = t.repo.CreateTiger(ctx, tiger, tigerSightings, email)
	if err != nil {
		fmt.Printf("error creating tiger: %+v", err)
		return customerr.GetTigersRepoError()
	}
	return nil
}

func (t *tiger) CreateSighting(ctx context.Context, createSightingReq *model.CreateTigerSightingRequest, email string) *customerr.APIError {
	parsedTime, err := time.Parse("2006-01-02T15:04:05Z", createSightingReq.LastSeen)
	if err != nil {
		return customerr.GetInvalidTimeError()
	}
	lat, err := strconv.ParseFloat(createSightingReq.Lat, 64)
	if err != nil {
		return &customerr.APIError{
			StatusCode: 400,
			Message:    "Invalid lat",
		}
	}
	lng, err := strconv.ParseFloat(createSightingReq.Lng, 64)
	if err != nil {
		return &customerr.APIError{
			StatusCode: 400,
			Message:    "Invalid lng",
		}
	}
	tigerSightings := &model.TigerSightings{
		LastSeen: parsedTime,
		Lat:      lat,
		Lng:      lng,
		TigerID:  createSightingReq.TigerId,
	}

	lastSighting, err := t.repo.GetLastSighting(ctx, createSightingReq.TigerId)
	if err != nil {
		fmt.Printf("error getting last sighting: %+v\n", err)
		return customerr.GetTigersRepoError()
	}
	if t.isDistantFromLastSighting(ctx, 5, lastSighting, tigerSightings) {
		err = t.repo.CreateSighting(ctx, nil, tigerSightings, email)
		if err != nil {
			fmt.Printf("error creating tiger sighting: %+v\n", err)
			return customerr.GetTigersRepoError()
		}
		return nil
	}
	return &customerr.APIError{
		StatusCode: 400,
		Message:    "Tiger is too close to last sighting",
	}
}

func (t *tiger) isDistantFromLastSighting(ctx context.Context, allowedDistance float64, lastSighting, currentSighting *model.TigerSightings) bool {
	radlat1 := float64(math.Pi * currentSighting.Lat / 180)
	radlat2 := float64(math.Pi * lastSighting.Lat / 180)

	theta := float64(currentSighting.Lng - currentSighting.Lng)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	dist = dist * 1.609344
	fmt.Println(dist)
	return dist > allowedDistance
}