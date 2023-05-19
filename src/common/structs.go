package common

import (
	sqlc "tienpvse1/go-fiber-server/src/modules/generated/sql"
)

type StructType int

type Bundler struct {
	Imports     []Bundler
	Controllers []IController
	Services    []interface{}
	Queries     *sqlc.Queries
	Router       *Router
}
