package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pranoyk/tiger-sightings/model"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	CreateTiger(ctx context.Context, tiger *model.Tiger) error
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CreateTiger(ctx context.Context, tiger *model.Tiger) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO tigers (name, dob, last_seen, lat, lng) VALUES ($1, $2, $3, $4, $5)", tiger.Name, tiger.Dob, tiger.LastSeen, tiger.Lat, tiger.Lng)
	if err != nil {
		fmt.Println("error inserting tiger: ", err)
		return err
	}
	return nil
}
