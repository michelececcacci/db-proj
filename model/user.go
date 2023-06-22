package model

import (
	"time"

	"github.com/michelececcacci/db-proj/queries"
)

func (m Model) InsertUser(arg queries.InsertUserParams, password string, signupDate time.Time) error {
	err := m.q.InsertUser(m.ctx, arg)
	if err != nil {
		return err
	}
	return m.q.InsertPassword(m.ctx, queries.InsertPasswordParams{
		Username:        arg.Username,
		Password:        password,
		Datainserimento: signupDate,
	})
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
		Username:      followed,
		Numeroseguaci: +1,
	})
	if err != nil {
		return err
	}
	return m.q.InsertFollower(m.ctx, arg)
}
