package view

// Subviews implement tea.Model and the main view is
// responsible for updating them and selecting the displayed one
// according to the state changes.

import (
	"context"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
	feed "github.com/michelececcacci/db-proj/view/feed"
	login "github.com/michelececcacci/db-proj/view/login"
	profile "github.com/michelececcacci/db-proj/view/profile"
	"github.com/michelececcacci/db-proj/view/signup"
)

type viewOption func(*mainView)

type state int

const (
	loginState state = iota
	profileState
	signupState
	feedState
	passwordResetState
)

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
	state             state
}

func NewMainView(options ...viewOption) mainView {
	m := mainView{}
	for _, opt := range options {
		opt(&m)
	}
	m.state = loginState // users need to login on startup
	m.loginView = login.New(&m.ctx, m.q)
	m.profileView = profile.New(&m.ctx, m.q, "user1") // TODO CHANGE
	m.signUpView = signup.New(&m.ctx, m.q)
	m.feedView = feed.New(&m.ctx, m.q, "user1") // TODO CHANGE
	m.help = newHelpComponent()
	m.passwordResetView = newPasswordResetView(&m.ctx, m.q, "user1") // TODO CHANGE
	return m
}

func (m mainView) Init() tea.Cmd {
	return nil
}

func (m mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.feedView, cmd = m.feedView.Update(msg) // always updated because of screen change messages
	}
	m.help, _ = m.help.Update(msg) // always updated
	switch m.state {
	case loginState:
		m.loginView, cmd = m.loginView.Update(msg)
		lv := m.loginView.(loginView)
		m.authUsername, m.authError = lv.GetAuthenticatedUsername()
	case profileState:
		m.profileView, _ = m.profileView.Update(msg)
	case signupState:
		m.signUpView, cmd = m.signUpView.Update(msg)
	case feedState:
		m.feedView, cmd = m.feedView.Update(msg)
	case passwordResetState:
		m.passwordResetView, cmd = m.passwordResetView.Update(msg)
	}
	return m, cmd
}

func (m mainView) View() string {
	var sb = strings.Builder{}
	switch m.state {
	case loginState:
		sb.WriteString(m.loginView.View())
	case profileState:
		sb.WriteString(m.profileView.View())
	case signupState:
		sb.WriteString(m.signUpView.View())
	case passwordResetState:
		sb.WriteString(m.passwordResetView.View())
	case feedState:
		sb.WriteString(m.feedView.View())
	}
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

type loginView interface {
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
	GetAuthenticatedUsername() (string, error)
}
