package model

import (
	"context"

	"github.com/michelececcacci/db-proj/queries"
)

func (m *Model) Authenticate(ctx context.Context, arg queries.AuthenticateParams) (int64, error) {
  return m.q.Authenticate(ctx, arg)
}


func (m *Model) GetPastPasswords(ctx context.Context, username string) ([]string, error) {
  return m.q.GetPastPasswords(ctx, username)
}

