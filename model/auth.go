package model

import (
	"context"
	"time"

	"github.com/michelececcacci/db-proj/queries"
)

func (m Model) Authenticate(arg queries.AuthenticateParams, loginDate time.Time) (bool, error) {
	r, err := m.q.Authenticate(m.ctx, arg)
	if r != 1 || err != nil {
		return false, err
	}
	return true, m.q.InsertUserAccess(m.ctx, queries.InsertUserAccessParams{
		Username:       arg.Username,
		Timestamplogin: loginDate,
	})
}

func (m *Model) GetPastPasswords(ctx context.Context, username string) ([]string, error) {
	return m.q.GetPastPasswords(ctx, username)
}
