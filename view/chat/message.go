package chat

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/michelececcacci/db-proj/queries"
)

type message struct {
	author    string
	content   string
	timestamp time.Time
}

func (m message) Title() string {
	return m.content // should be limited in some way
}

func (m message) FilterValue() string {
	return m.author + "\n" + m.timestamp.String() + "\n" + m.content
}

func (m message) Description() string {
	return m.content
}

func messagesToItems(msgs []queries.GetChatMessagesRow) []list.Item {
	var l []list.Item
	for _, msg := range msgs {
		l = append(l, message{
			author:    msg.Username,
			content:   msg.Testo,
			timestamp: msg.Timestampinvio,
		})
	}
	return l
}
