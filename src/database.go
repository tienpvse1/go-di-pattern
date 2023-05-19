package database

import (
	"database/sql"
	sqlc "tienpvse1/go-fiber-server/src/modules/generated/sql"
)

func InitDatabase() (*sqlc.Queries, error){
  db, err := sql.Open("postgres", "user=postgres dbname=go password=password sslmode=disable")
	if err != nil {
		return nil, err
	}

	queries := sqlc.New(db)
  return queries, nil
}

