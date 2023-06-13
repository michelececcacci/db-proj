package profile

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/util"
)

type state int

const (
	followingState = iota
	followersState
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
	state
}

func (p profileView) View() string {
	if p.state == followersState {
		return util.ListStyle.Render(p.followers.View())
	}
	return util.ListStyle.Render(p.following.View())
}

func (p profileView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return p, tea.Quit
		case tea.KeyLeft, tea.KeyRight:
			if p.state == followingState {
				p.state = followersState
			} else if p.state == followersState {
				p.state = followingState
			}
		case tea.KeyEnter:
			var selected list.Item
			if p.state == followersState {
				selected = p.followers.SelectedItem()
			} else if p.state == followingState {
				selected = p.following.SelectedItem()
			}
			if selected != nil {
				return New(p.ctx, p.q, selected.FilterValue()), nil
			}
		}
	case tea.WindowSizeMsg:
	}
	if p.state == followersState {
		p.followers, cmd = p.followers.Update(msg)
	} else if p.state == followingState {
		p.following, cmd = p.following.Update(msg)
	}
	return p, cmd
}

func (p profileView) Init() tea.Cmd {
	return nil
}

func New(ctx *context.Context, q *queries.Queries, username string) profileView {
	following, _ := q.GetFollowing(*ctx, username)
	followers, _ := q.GetFollowers(*ctx, username)
	p := profileView{
		username:  username,
		location:  "test_location", // TODO CHANGE
		followers: list.New(toUser(followers), list.NewDefaultDelegate(), 25, 25),
		following: list.New(toUser(following), list.NewDefaultDelegate(), 25, 25),
		ctx:       ctx,
		q:         q,
		state:     followersState,
	}
	p.followers.Title = fmt.Sprintf("Followed by %s", p.username)
	p.following.Title = fmt.Sprintf("Following %s", p.username)
	return p
}

func toUser(usernames []string) []list.Item {
	var u []list.Item
	for _, s := range usernames {
		u = append(u, user{username: s})
	}
	return u
}
