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

func (m Model) ExitMember(memberId int32, exitDate time.Time, removerAdmin sql.NullInt32, motivation sql.NullString) error {
	// check if the user is member of the chat
	nullableJoinDate, err := m.q.CheckIfMemberStillInChat(m.ctx, memberId)
	if err != nil {
		return err
	}
	if !nullableJoinDate.Valid {
		return errors.New("The user is not member of the chat\n")
	}
	joinDate := nullableJoinDate.Time
	if exitDate.Before(joinDate) {
		return errors.New("The exit date of the member is before join date")
	}
	return m.q.InsertExit(m.ctx, queries.InsertExitParams{
		Idmembro:             memberId,
		Datauscita:           exitDate,
		Motivazione:          motivation,
		Idmembroresponsabile: removerAdmin,
	})
}

func (m Model) ExileMember(memberId int32, exitDate time.Time, removerAdminId int32, motivation sql.NullString) error {
	return m.ExitMember(memberId, exitDate, sql.NullInt32{Int32: removerAdminId, Valid: true}, motivation)
}

func (m Model) IntentionalExitMember(memberId int32, exitDate time.Time, motivation sql.NullString) error {
	return m.ExitMember(memberId, exitDate, sql.NullInt32{Int32: 0, Valid: false}, motivation)
}

func (m Model) GetMemberData(memberId int32) (queries.GetDataOfMemberRow, error) {
	return m.q.GetDataOfMember(m.ctx, memberId)
}

func (m Model) CheckIfUserStillInChat(username string, idchat int32) (bool, error) {
	count, err := m.q.CheckIfUserStillInChat(m.ctx, queries.CheckIfUserStillInChatParams{
		Username: username,
		Idchat:   idchat,
	})
	return count == 1, err
}

func (m Model) GetFullFeed(username string) ([]queries.GetFullFeedRow, error) {
	return m.q.GetFullFeed(m.ctx, username)
}

func (m Model) GetCurrentChats(username string) ([]int32, error) {
	var currentIds []int32
	res, err := m.q.GetChatIds(m.ctx, username)
	if err != nil {
		return nil, err
	}
	for _, r := range res {
		ok, err := m.CheckIfUserStillInChat(username, r.Idchat)
		if err != nil {
			return nil, err
		}
		if ok {
			currentIds = append(currentIds, r.Idchat)
		}
	}
	return currentIds, nil
}

func (m Model) GetChatInfos(id int32) (queries.GetChatInfosRow, error) {
	return m.q.GetChatInfos(m.ctx, id)
}

func (m Model) InsertMessage(message queries.InsertMessageParams) error {
	return m.q.InsertMessage(m.ctx, message)
}

func (m Model) GetChatMessages(idChat int32) ([]queries.GetChatMessagesRow, error) {
	return m.q.GetChatMessages(m.ctx, idChat)
}

func (m Model) GetChatMembers(id int32) ([]int32, error) {
	return m.q.GetChatMembers(m.ctx, id)
}
