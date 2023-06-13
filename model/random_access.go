package model

import (
	"database/sql"

	"github.com/michelececcacci/db-proj/queries"
)

func (m *Model) GetRandomUser() (string, error) {
	return m.q.GetRandomUser(m.ctx)
}

func (m Model) GetRandomChat() (int32, error) {
	return m.q.GetRandomChat(m.ctx)
}

func (m Model) GetRandomAdminInChat(idChat int32) (int32, error) {
	return m.q.GetRandomAdminInChat(m.ctx, idChat)
}

func (m Model) GetAllPossibleMembers() ([]queries.GetAllPossibleMembersRow, error) {
	return m.q.GetAllPossibleMembers(m.ctx)
}

func (m Model) GetAllPossibleFollowings() ([]queries.GetAllPossibleFollowingsRow, error) {
	return m.q.GetAllPossibleFollowings(m.ctx)
}

func (m Model) GetRandomCurrentMemberInChat(chat int32) (sql.NullInt32, sql.NullTime, error) {
	s, err := m.q.GetRandomMemberInChat(m.ctx, chat)
	return s.Idmembro, s.Dataentrata, err
}
