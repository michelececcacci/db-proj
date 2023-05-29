package model

func (m *Model) GetRandomUser() (string, error) {
   return m.q.GetRandomUser(m.ctx)
}

func (m Model) GetRandomChat() (int32, error) {
  return m.q.GetRandomChat(m.ctx)
}

func (m Model) GetRandomAdminInChat(idChat int32) (int32, error) {
  return m.q.GetRandomAdminInChat(m.ctx, idChat)
}

