package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type Sneaker struct {
	ID          pgtype.UUID `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Created_at  time.Time   `json:"created_at"`
}
