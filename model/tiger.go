package model

import (
	"time"
)

type CreateTigerRequest struct {
	Name     string `json:"name" binding:"required"`
	Dob      string `json:"dob" binding:"required"`
	LastSeen string `json:"last_seen" binding:"required"`
	Lat      string `json:"lat" binding:"required"`
	Lng      string `json:"lng" binding:"required"`
}

type CreateTigerSightingRequest struct {
	LastSeen string `json:"last_seen" binding:"required"`
	Lat      string `json:"lat" binding:"required"`
	Lng      string `json:"lng" binding:"required"`
	TigerId  string `json:"tiger_id"`
}

type Tiger struct {
	ID        string    `json:"id" db:"tiger_id"`
	Name      string    `json:"name" db:"name"`
	Dob       time.Time `json:"dob" db:"dob"`
	CreatedBy string    `json:"created_by" db:"created_by"`
}

type TigerSightings struct {
	ID         string    `json:"id" db:"tiger_sightings_id"`
	TigerID    string    `json:"tiger_id" db:"tiger_id"`
	CreatedBy  string    `json:"created_by" db:"created_by"`
	Created_At string    `json:"created_at" db:"created_at"`
	LastSeen   time.Time `json:"last_seen" db:"last_seen"`
	Lat        float64   `json:"lat" db:"lat"`
	Lng        float64   `json:"lng" db:"lng"`
}
