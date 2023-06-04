package feed

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/michelececcacci/db-proj/queries"
)

type post struct {
	author    string
	content   string
	title     string
	upvotes   int
	timestamp time.Time
}

func (p post) Title() string {
	return p.title
}

func (p post) FilterValue() string {
	return p.title + p.author
}

func (p post) Description() string {
	return fmt.Sprintf("Author: %s Title: %s", p.author, p.title)
}

func toPost(result []queries.GetFullFeedRow) []list.Item {
	var posts []list.Item
	for _, p := range result {
		var title string
		if p.Titolo.Valid {
			title = p.Titolo.String
		}
		posts = append(posts, post{
			author:  p.Autore,
			title:   title,
			content: p.Testo,
			upvotes: 0,
		})
	}
	return posts
}
