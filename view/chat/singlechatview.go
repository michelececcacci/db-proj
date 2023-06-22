package chat

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
)

type state int

const (
	viewMessage state = iota
	sendMessage
)

type singleChatView struct {
	username     string
	name         string
	messages     list.Model
	messageInput textinput.Model
	state        state
	model        chatModel
	chatInfo     chatInfo
}

func (c singleChatView) Init() tea.Cmd { return nil }

func (c singleChatView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch c.state {
	case viewMessage:
		c.messages, cmd = c.messages.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				c.state = sendMessage
				c.messageInput.Focus()
			}
		}
	case sendMessage:
		c.messageInput, cmd = c.messageInput.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				c.sendMessage()
				return newSingleChatView(c.chatInfo, c.username, c.model), nil
			}
		}
	}
	return c, cmd
}

func (c singleChatView) sendMessage() {
	id, _ := c.model.GetChatUserId(c.chatInfo.id, c.username)
	{
		c.model.InsertMessage(queries.InsertMessageParams{
			Testo:          c.messageInput.Value(),
			Mittente:       id,
			Timestampinvio: time.Now().UTC(),
		})
	}
}

func (c singleChatView) View() string {
	var sb strings.Builder
	sb.WriteString(c.name + "\n")
	switch c.state {
	case viewMessage:
		sb.WriteString(c.messages.View())
	case sendMessage:
		sb.WriteString(c.messageInput.View())
	}
	return sb.String()
}

func newSingleChatView(ci chatInfo, username string, m chatModel) singleChatView {
	info, _ := m.GetChatMessages(ci.id)
	c := singleChatView{
		model:        m,
		messageInput: textinput.New(),
		messages:     list.New(messagesToItems(info), list.NewDefaultDelegate(), 40, 25),
		state:        viewMessage,
		username:     username,
		chatInfo:     ci,
	}
	c.messages.Title = ci.name
	return c
}
