package model

import "github.com/michelececcacci/db-proj/queries"

func (m Model) InsertUser(arg queries.InsertUserParams) error {
	return m.q.InsertUser(m.ctx, arg)
}

func (m Model) GetFollowers(usernameseguito string) ([]string, error) {
	return m.q.GetFollowers(m.ctx, usernameseguito)
}

func (m Model) GetFollowing(usernameseguace string) ([]string, error) {
	return m.q.GetFollowing(m.ctx, usernameseguace)
}

func (m Model) InsertFollower(arg queries.InsertFollowerParams) error {
	followed := arg.Usernameseguito
	err := m.q.UpdateNumberOfFollowers(m.ctx, queries.UpdateNumberOfFollowersParams{
		Username:     followed, 
		Numeroseguaci: +1,
	})
	if err != nil {
		return err
	}
	return m.q.InsertFollower(m.ctx, arg)
}
