package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Item struct {
	Task   string
	Status string
}

type DB struct {
	pool *pgxpool.Pool
}

func New(user, password, dbname, host string, port int) (*DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the database: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("Failed to ping the database: %w", err)
	}
	return &DB{pool: pool}, nil
}

func (db *DB) InsertItem(ctx context.Context, item Item) error {
	query := `INSERT INTO todo_items(task, status) VALUES ($1, $2)`
	_, err := db.pool.Exec(ctx, query, item.Task, item.Status)
	return err
}
func (db *DB) GetAllItems(ctx context.Context) ([]Item, error) {
	query := `SELECT task, status FROM todo_items`
	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch all items: %w", err)
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.Task, &item.Status)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan item: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Failed to iterate rows: %w", err)
	}

	return items, nil
}

func (db *DB) Close() {
	db.pool.Close()
}
