package model

import "github.com/michelececcacci/db-proj/queries"

func (m Model) GetFullFeed(username string) ([]queries.GetFullFeedRow, error) {
	return m.q.GetFullFeed(m.ctx, username)
}

func (m Model) InsertPost(arg queries.InsertPostParams) error {
	return m.q.InsertPost(m.ctx, arg)
}

func (m Model) InsertComment(arg queries.InsertCommentParams) error {
	return m.q.InsertComment(m.ctx, arg)
}
