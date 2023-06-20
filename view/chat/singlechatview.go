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
	username   string
	name       string
	messages   list.Model
	sendMessage textinput.Model
	state      state
	model      chatModel
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
				{
					c.model.InsertMessage(queries.InsertMessageParams{
						Testo:          c.sendMessage.Value(),
						Mittente:       0, // todo fix
						Timestampinvio: time.Now().UTC(),
					})
				}
			}
		}
	case sendMessage:
		c.sendMessage, cmd = c.sendMessage.Update(msg)
	}
	return c, cmd
}

func (c singleChatView) View() string {
	var sb strings.Builder
	sb.WriteString(c.name + "\n")
	switch c.state {
	case viewMessage:
		sb.WriteString(c.messages.View())
	case sendMessage:
		sb.WriteString(c.sendMessage.View())
	}
	return sb.String()
}

func newSingleChatView(ci chatInfo, username string, m chatModel) singleChatView {
	info, _ := m.GetChatMessages(ci.id)
	c := singleChatView{
		model:      m,
		sendMessage: textinput.New(),
		messages:   list.New(messagesToItems(info), list.NewDefaultDelegate(), 40, 25),
		state:      viewMessage,
		username:   username,
	}
	c.messages.Title = ci.name
	return c
}
