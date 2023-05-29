package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"time"

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

func (m *Model) GetFollowers(usernameseguito string) ([]string, error) {
	return m.q.GetFollowers(m.ctx, usernameseguito)
}

func (m *Model) GetFollowing(usernameseguace string) ([]string, error) {
	return m.q.GetFollowing(m.ctx, usernameseguace)
}

func (m *Model) InsertUser(arg queries.InsertUserParams) error {
	return m.q.InsertUser(m.ctx, arg) 
} 

// create a new chat and create the first member of the chat
func (m *Model) InsertChat(arg queries.InsertChatParams, creator string) error {
	idChat, err := m.q.InsertChat(m.ctx, arg)
  if err != nil {
    return err
  }
  memberId, err := m.q.InsertMember(m.ctx, queries.InsertMemberParams{
		Dataentrata: time.Now(),
		Username:    creator,
		Idchat:      idChat,
		Amministratore: sql.NullInt32{
			Int32: 0,
			Valid: false,
		},
	})
  return m.q.InsertAdmin(m.ctx, memberId)
}

func (m *Model) InsertFollower(arg queries.InsertFollowerParams) error {
	return m.q.InsertFollower(m.ctx, arg)
}

type InsertMemberParams struct {
	Username string
	Idchat   int32
}

func (m *Model) IsValidAdmin(admin int32, idChat int32) (bool, error) {
	count, err := m.q.IsValidAdmin(m.ctx, queries.IsValidAdminParams{
		Idmembro: admin,
		Idchat:   idChat,
	})
	return count == 1, err
}

// create a new member who is not admin
func (m *Model) InsertNotAdminMember(username string, idChat int32, inserterAdmin int32) (int32, error) {
	count, err := m.q.CheckIfUserStillInChat(m.ctx, queries.CheckIfUserStillInChatParams{Username: username, Idchat: idChat})
	if err != nil {
		return 0, err
	}
	if count != 0 {
		return 0, errors.New(fmt.Sprintf("The user %s is still member of chat %d", username, idChat))
	}
	member := queries.InsertMemberParams{
			Dataentrata: time.Now(),
			Username:    username,
			Idchat:      idChat,
			Amministratore: sql.NullInt32{
				Int32: inserterAdmin,
				Valid: true,
			},
		}
	isValidAdmin, err := m.IsValidAdmin(inserterAdmin, idChat)
	if isValidAdmin && err == nil {
		return m.q.InsertMember(m.ctx, member)
	} else {
		return 0, errors.New("The inserter admin is not an admin of the chat")
	}
}

// create an admin
func (m *Model) InsertAdminMember(username string, idChat int32, inserterAdmin int32) (int32, error) {
	memberId, err := m.InsertNotAdminMember(username, idChat, inserterAdmin)
  if err != nil {
		return memberId, err
	}
	return memberId, m.q.InsertAdmin(m.ctx, memberId)
}
