package feed

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/michelececcacci/db-proj/queries"
)

type post struct {
	Author    string
	Content   string
	PostTitle string
	upvotes   int
	timestamp time.Time
	id int32
}

func (p post) Title() string {
	return p.PostTitle
}

func (p post) FilterValue() string {
	return p.PostTitle + p.Author
}

func (p post) Description() string {
	return fmt.Sprintf("Author: %s Title: %s", p.Author, p.PostTitle)
}

func toPost(result []queries.GetFullFeedRow) []list.Item {
	var posts []list.Item
	for _, p := range result {
		var title string
		if p.Titolo.Valid {
			title = p.Titolo.String
		}
		posts = append(posts, post{
			Author:    p.Autore,
			PostTitle: title,
			Content:   p.Testo,
			upvotes:   0,
			id: p.Idcontenuto,
		})
	}
	return posts
}
