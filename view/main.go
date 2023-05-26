package view

// Subviews implement tea.Model and the main view is
// responsible for updating them and selecting the displayed one
// according to the state changes.

import (
	"context"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/view/components"
)

type viewOption func(*mainView)

type mainView struct {
	ctx          context.Context
	q            *queries.Queries
	loginView    tea.Model
	profileView  tea.Model
	registerView tea.Model
	feedView     tea.Model
	help         tea.Model
}

func NewMainView(options ...viewOption) mainView {
	m := mainView{}
	for _, opt := range options {
		opt(&m)
	}
	m.loginView = newLoginView(&m.ctx, m.q)
	m.profileView = newProfileView(&m.ctx, m.q, "user1")
	m.registerView = newRegisterView(&m.ctx, m.q)
	m.feedView = NewFeedView(&m.ctx, m.q)
	m.help = components.NewHelpComponent()
	return m
}

func (m mainView) Init() tea.Cmd {
	return nil
}

func (m mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.help, _ = m.help.Update(msg) // always updated
	m.loginView, cmd = m.loginView.Update(msg)
	// return m.loginView.Update(msg)
	m.profileView, _ = m.profileView.Update(msg)
	// m.registerView, cmd = m.registerView.Update(msg)
	m.feedView, cmd = m.feedView.Update(msg)
	return m, cmd
}

func (m mainView) View() string {
	var sb = strings.Builder{}
	sb.WriteString(m.profileView.View())
	// sb.WriteString(m.registerView.View())
	// sb.WriteString(m.feedView.View())
	sb.WriteString(m.help.View())
	return sb.String()
}

// ctx is needed to query the database
func WithContext(ctx context.Context) viewOption {
	return func(m *mainView) {
		m.ctx = ctx
	}
}

// q is needed to query the database
func WithQueries(q *queries.Queries) viewOption {
	return func(m *mainView) {
		m.q = q
	}
}
