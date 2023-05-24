package user

import (
	"context"
	sqlc "tienpvse1/go-fiber-server/src/modules/generated/sql"
)

type UserService struct {
	Queries *sqlc.Queries `inject:"sqlc_queries"`
	Context context.Context
}

func (service UserService) FindOne() ([]sqlc.Author, error) {
	return service.Queries.ListAuthors(context.Background())
}
