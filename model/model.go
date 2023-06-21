package model

import (
	"context"

	"github.com/michelececcacci/db-proj/queries"
)

type Model struct {
	q   *queries.Queries
	ctx context.Context
}

func New(db queries.DBTX, ctx context.Context) Model {
	return Model{
		q:   queries.New(db),
		ctx: ctx,
	}
}



