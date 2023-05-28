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
	"github.com/michelececcacci/db-proj/view/signup"
)

type viewOption func(*mainView)

type mainView struct {
	ctx               context.Context
	q                 *queries.Queries
	loginView         tea.Model
	profileView       tea.Model
	signUpView        tea.Model
	feedView          tea.Model
	help              tea.Model
	passwordResetView tea.Model
}

func NewMainView(options ...viewOption) mainView {
	m := mainView{}
	for _, opt := range options {
		opt(&m)
	}
	m.loginView = newLoginView(&m.ctx, m.q)
	m.profileView = newProfileView(&m.ctx, m.q, "user1")
	m.signUpView = signup.New(&m.ctx, m.q)
	m.feedView = NewFeedView(&m.ctx, m.q)
	m.help = components.NewHelpComponent()
	m.passwordResetView = newPasswordResetView(&m.ctx, m.q, "user1") // TODO CHANGE
	return m
}

func (m mainView) Init() tea.Cmd {
	return nil
}

func (m mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// TODO we don't need to update everything, but this is pretty useful for testing
	var cmd tea.Cmd
	m.help, cmd = m.help.Update(msg) // always updated
	// m.loginView, cmd = m.loginView.Update(msg)
	// m.profileView, _ = m.profileView.Update(msg)
	m.signUpView, cmd = m.signUpView.Update(msg)
	// m.feedView, cmd = m.feedView.Update(msg)
	// m.passwordResetView, cmd = m.passwordResetView.Update(msg)
	return m, cmd
}

func (m mainView) View() string {
	var sb = strings.Builder{}
	// sb.WriteString(m.profileView.View())
	// sb.WriteString(m.loginView.View())
	sb.WriteString(m.signUpView.View())
	// sb.WriteString(m.passwordResetView.View())
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
