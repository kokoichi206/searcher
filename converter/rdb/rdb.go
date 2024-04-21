package rdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type database struct {
	db *sql.DB
}

func New() (*database, error) {
	// 固定値を使っている。
	source := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		"localhost", "5445", "ubuntu", "sakamichi", "nogi-official", "disable",
	)

	sqlDB, err := sql.Open("postgres", source)
	if err != nil {
		return nil, fmt.Errorf("failed to open sql: %w", err)
	}

	db := &database{
		db: sqlDB,
	}

	return db, nil
}
