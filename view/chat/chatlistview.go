package chat

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
)

const (
	multipleChatsState state = iota
	singleChatState
)

type chatlistView struct {
	list     list.Model
	chatView tea.Model
	state    state
	username string
	model    chatModel
}

func (c chatlistView) Init() tea.Cmd { return nil }

func (c chatlistView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if c.state == multipleChatsState {
				c.switchToChat()
			}
		case tea.KeyCtrlC:
			return c, tea.Quit
		}
	}
	switch c.state {
	case singleChatState:
		c.chatView, cmd = c.chatView.Update(msg)
	case multipleChatsState:
		c.list, cmd = c.list.Update(msg)
	}
	return c, cmd
}

func (c *chatlistView) switchToChat() {
	item := c.list.SelectedItem()
	if item != nil {
		info := item.(chatInfo)
		c.chatView = newSingleChatView(info, c.username, c.model)
		c.state = singleChatState
	}
}

func (c chatlistView) View() string {
	var sb strings.Builder
	switch c.state {
	case singleChatState:
		sb.WriteString(c.chatView.View()) // this will panic if the state is illegal, just as intended
	case multipleChatsState:
		sb.WriteString(c.list.View())
	}
	return sb.String()
}

func NewChatListView(username string, cm chatModel) chatlistView {
	ids, _ := cm.GetCurrentChats(username)
	var items []list.Item
	for _, id := range ids {
		info, _ := cm.GetChatInfos(id)
		items = append(items, chatInfo{
			name:        info.Nome,
			id:          id,
			description: info.Descrizione,
		})
	}
	c := chatlistView{
		list:     list.New(items, list.NewDefaultDelegate(), 30, 30),
		chatView: nil,
		state:    multipleChatsState,
		username: username,
		model:    cm,
	}
	c.list.Title = fmt.Sprintf("%s's chats", username)
	return c
}

type chatModel interface {
	GetCurrentChats(string) ([]int32, error)
	GetChatInfos(int32) (queries.GetChatInfosRow, error)
	InsertMessage(queries.InsertMessageParams) error
	GetChatMessages(idChat int32) ([]queries.GetChatMessagesRow, error)
}
