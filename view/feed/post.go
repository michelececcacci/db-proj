package feed

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/michelececcacci/db-proj/queries"
)


type post struct{
	author string
	content string
	title string
	upvotes int
	timestamp time.Time
}

func (p post) Title() string {
	return p.title
}

func (p post) FilterValue() string {
	return p.Title()
}

func (p post) Description() string {
	return fmt.Sprintf("Author: %s, Posted at: %s", p.author, p.timestamp.String())
}

func f(result []queries.GetFullFeedRow) []list.Item {
	return nil
}