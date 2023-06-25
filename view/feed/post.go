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
	id        int32
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
		posts = append(posts, post{
			Author:    p.Autore,
			PostTitle: p.Titolo.String, // post has always a title
			Content:   p.Testo,
			upvotes:   int(p.Likedelta),
			id:        p.Idcontenuto,
			timestamp: p.Timestamppubblicazione,
		})
	}
	return posts
}
