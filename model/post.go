package model

import "github.com/michelececcacci/db-proj/queries"

func (m Model) GetFullFeed(username string) ([]queries.GetFullFeedRow, error) {
	return m.q.GetFullFeed(m.ctx, username)
}

