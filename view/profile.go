package view

import (
	"context"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/util"

	tea "github.com/charmbracelet/bubbletea"
)

type user struct {
	username string
	location string
}

func (u user) Title() string {
	return u.username
}

func (u user) Description() string {
	return ""
}

func (u user) FilterValue() string {
	return u.Title()
}

// Read only view for users. Still WIP.
type profileView struct {
	username  string
	location  string
	followers list.Model
	following list.Model
	current   int
	ctx       *context.Context
	q         *queries.Queries
}

func (p profileView) View() string {
	var sb strings.Builder
	sb.WriteString(p.followers.View() + "\n")
	sb.WriteString(p.following.View() + "\n")
	return sb.String()
}

func (p profileView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return p, tea.Quit
		case tea.KeyUp, tea.KeyTab:
			p.current = util.Max(0, p.current-1)
		case tea.KeyDown, tea.KeyShiftTab:
			p.current = util.Min(1, p.current+1)

		}
	case tea.WindowSizeMsg:
	}
	p.followers, cmd = p.followers.Update(msg)
	p.following, cmd = p.following.Update(msg)
	return p, cmd
}

func (p profileView) Init() tea.Cmd {
	return nil
}

func newProfileView(ctx *context.Context, q *queries.Queries, username string) profileView {
	following , _:= q.GetFollowing(*ctx, username)
	followers, _ := q.GetFollowers(*ctx, username)
	return profileView{
		username:  "test_username",
		location:  "test_location",
		followers: list.New(toUser(followers), list.NewDefaultDelegate(), 10, 10),
		following: list.New(toUser(following), list.NewDefaultDelegate(), 10, 10),
		ctx:       ctx,
		q:         q,
	}
}

func toUser(usernames []string) []list.Item {
	var u []list.Item
	for _,s  := range usernames {
		u = append(u, user{username: s})
	}
	return u
}
