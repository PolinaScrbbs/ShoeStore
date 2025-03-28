package queries

import (
	"ShoeStore/database/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllSneakers(pool *pgxpool.Pool) ([]models.Sneaker, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT id, title, description, price, created_at FROM sneakers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sneakers []models.Sneaker
	for rows.Next() {
		var sneaker models.Sneaker
		if err := rows.Scan(
			&sneaker.ID, &sneaker.Title, &sneaker.Description, &sneaker.Price, &sneaker.Created_at); err != nil {
			return nil, err
		}
		sneakers = append(sneakers, sneaker)
	}
	return sneakers, nil
}
