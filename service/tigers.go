package service

// go:generate mockgen -source=./service/tigers.go -destination=./mocks/mock_tigers.go -package=mocks github.com/pranoyk/tiger-sightings/service Tiger

import (
	"context"
	"fmt"
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
		return customerr.GetCreateTigerRepoError()
	}
	return nil
}
