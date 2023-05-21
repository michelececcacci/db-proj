package view

import (
	"context"
	"fmt"
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
	sb.WriteString(fmt.Sprintf("Username: %s\n", p.username))
	sb.WriteString(fmt.Sprintf("Location: %s\n", p.location))
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
			// p.current = util.Min(len(p)-1)

		}
	case tea.WindowSizeMsg:
	}
	p.followers, _ = p.followers.Update(msg)
	p.following, cmd = p.following.Update(msg)
	return p, cmd
}

func (p profileView) Init() tea.Cmd {
	return nil
}

func newProfileView(ctx *context.Context, q *queries.Queries) profileView {
	following := []list.Item{user{username: "user_1"}}
	followers := []list.Item{user{username: "user_2"}}
	return profileView{
		username:  "test_username",
		location:  "test_location",
		followers: list.New(followers, list.NewDefaultDelegate(), 0, 0),
		following: list.New(following, list.NewDefaultDelegate(), 0, 0),
		ctx:       ctx,
		q:         q,
	}
}
