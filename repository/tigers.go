package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pranoyk/tiger-sightings/model"
)

type tigersRepository struct {
	db *sql.DB
}

type TigersRepository interface {
	CreateTiger(ctx context.Context, tiger *model.Tiger, tigerSightins *model.TigerSightings, email string) error
	CreateSighting(ctx context.Context, tx *sql.Tx, tigerSighting *model.TigerSightings, email string) (err error)
	GetLastSighting(ctx context.Context, tigerID string) (*model.TigerSightings, error)
	GetTigers(ctx context.Context) ([]*model.Tiger, error)
}

func NewTigersRepository(db *sql.DB) TigersRepository {
	return &tigersRepository{
		db: db,
	}
}

func (r *tigersRepository) GetTigers(ctx context.Context) ([]*model.Tiger, error) {
	var tigers []*model.Tiger
	rows, err := r.db.QueryContext(ctx, "SELECT t.name, t.dob, ts.last_seen, ts.lat, ts.lng FROM tigers t JOIN tiger_sightings ts ON t.tiger_id = ts.tiger_id ORDER BY ts.last_seen DESC")
	if err != nil {
		fmt.Println("error getting tigers: ", err)
		return nil, err
	}
	for rows.Next() {
		var tiger model.Tiger
		err = rows.Scan(&tiger.Name, &tiger.Dob, &tiger.LastSeen, &tiger.Lat, &tiger.Lng)
		if err != nil {
			fmt.Println("error scanning tigers: ", err)
			return nil, err
		}
		tigers = append(tigers, &tiger)
	}
	return tigers, nil
}

func (r *tigersRepository) CreateTiger(ctx context.Context, tiger *model.Tiger, tigerSightings *model.TigerSightings, email string) (err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				fmt.Println("error rolling back transaction: ", rollbackErr)
			}
			return
		}
		if commitErr := tx.Commit(); commitErr != nil {
			fmt.Println("error committing transaction: ", commitErr)
			return
		}
	}()
	_, err = tx.ExecContext(ctx, "INSERT INTO tigers (name, dob, created_by) VALUES ($1, $2, (SELECT user_id FROM users WHERE email = $3))", tiger.Name, tiger.Dob, email)
	if err != nil {
		fmt.Println("error inserting tiger: ", err)
		return err
	}

	err = tx.QueryRowContext(ctx, "SELECT tiger_id FROM tigers WHERE name = $1", tiger.Name).Scan(&tigerSightings.TigerID)
	if err != nil {
		return err
	}

	err = r.CreateSighting(ctx, tx, tigerSightings, email)
	if err != nil {
		return err
	}
	return nil
}

func (r *tigersRepository) CreateSighting(ctx context.Context, tx *sql.Tx, tigerSighting *model.TigerSightings, email string) (err error) {
	if tx != nil {
		_, err = tx.ExecContext(ctx, "INSERT INTO tiger_sightings (tiger_id, created_by, last_seen, lat, lng) VALUES ($1, (SELECT user_id FROM users WHERE email = $2), $3, $4, $5)",
			tigerSighting.TigerID,
			email, tigerSighting.LastSeen,
			tigerSighting.Lat,
			tigerSighting.Lng,
		)
	} else {
		fmt.Printf("tiger sighting: %+v\n", tigerSighting)
		fmt.Printf("email: %+v\n", email)
		_, err = r.db.ExecContext(ctx, "INSERT INTO tiger_sightings (tiger_id, created_by, last_seen, lat, lng) VALUES ($1, (SELECT user_id FROM users WHERE email = $2), $3, $4, $5)",
			tigerSighting.TigerID,
			email, tigerSighting.LastSeen,
			tigerSighting.Lat,
			tigerSighting.Lng,
		)
	}
	if err != nil {
		fmt.Println("error inserting tiger sighting: ", err)
		return err
	}
	return nil
}

func (r *tigersRepository) GetLastSighting(ctx context.Context, tigerID string) (*model.TigerSightings, error) {
	var tigerSighting model.TigerSightings
	err := r.db.QueryRowContext(ctx, "SELECT last_seen, lat, lng FROM tiger_sightings WHERE tiger_id = $1 ORDER BY last_seen DESC LIMIT 1", tigerID).Scan(&tigerSighting.LastSeen, &tigerSighting.Lat, &tigerSighting.Lng)
	if err != nil {
		fmt.Println("error getting last sighting: ", err)
		return nil, err
	}
	return &tigerSighting, nil
}
