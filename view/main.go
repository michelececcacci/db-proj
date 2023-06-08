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
	feed "github.com/michelececcacci/db-proj/view/feed"
	login "github.com/michelececcacci/db-proj/view/login"
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
	authUsername      string
	authError         error
}

func NewMainView(options ...viewOption) mainView {
	m := mainView{}
	for _, opt := range options {
		opt(&m)
	}
	m.loginView = login.New(&m.ctx, m.q)
	m.profileView = newProfileView(&m.ctx, m.q, "user1") // TODO CHANGE
	m.signUpView = signup.New(&m.ctx, m.q)
	m.feedView = feed.New(&m.ctx, m.q, "user1") // TODO CHANGE
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
	m.help, _ = m.help.Update(msg) // always updated
	// m.loginView, cmd = m.loginView.Update(msg)
	lv := m.loginView.(loginView)
	m.authUsername, m.authError = lv.GetAuthenticatedUsername()
	// m.profileView, _ = m.profileView.Update(msg)
	// m.signUpView, cmd = m.signUpView.Update(msg)
	m.feedView, cmd = m.feedView.Update(msg) // always updated because of screen change messages
	// m.passwordResetView, cmd = m.passwordResetView.Update(msg)
	return m, cmd
}

func (m mainView) View() string {
	var sb = strings.Builder{}
	// sb.WriteString(m.profileView.View())
	// sb.WriteString(m.loginView.View())
	// sb.WriteString(m.signUpView.View())
	// sb.WriteString(m.passwordResetView.View())
	sb.WriteString(m.feedView.View())
	// sb.WriteString(m.help.View())
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

type loginView interface {
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
	GetAuthenticatedUsername() (string, error)
}
