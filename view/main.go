package view

// Subviews implement tea.Model and the main view is
// responsible for updating them and selecting the displayed one
// according to the state changes.

import (
	"context"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/model"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/util"
	chat "github.com/michelececcacci/db-proj/view/chat"
	feed "github.com/michelececcacci/db-proj/view/feed"
	login "github.com/michelececcacci/db-proj/view/login"
	"github.com/michelececcacci/db-proj/view/post"
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
	chatState
	publishPostState
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
	chatView          tea.Model
	publishPostView   tea.Model
	authUsername      *string
	state             state
	model             *model.Model
}

func NewMainView(options ...viewOption) mainView {
	m := mainView{}
	for _, opt := range options {
		opt(&m)
	}
	m.state = loginState // users need to login on startup
	m.loginView = login.New(m.model)
	m.profileView = profile.New(m.model, m.authUsername)
	m.signUpView = signup.New(&m.ctx, m.q)
	m.feedView = feed.New(m.model, m.model, m.authUsername)
	m.help = newHelpComponent()
	m.passwordResetView = newPasswordResetView(&m.ctx, m.q, m.authUsername)
	m.chatView = chat.NewChatListView(m.authUsername, m.model)
	m.publishPostView = post.New(m.authUsername, m.model)
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
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlL:
			if m.state != signupState { // currently ctrl + l is used as the location selection shortcut in
				// the signup view, so we can't just update this state in that case.
				m.state = loginState
			}
		case tea.KeyCtrlS:
			m.state = signupState
		case tea.KeyCtrlF:
			m.state = feedState
		case tea.KeyCtrlA:
			m.state = chatState
		case tea.KeyCtrlW:
			m.state = publishPostState
		}
	}
	m.help, _ = m.help.Update(msg) // always updated
	switch m.state {
	case loginState:
		m.loginView, cmd = m.loginView.Update(msg)
		lv := m.loginView.(loginView)
		username, err := lv.GetAuthenticatedUsername()
		if err == nil {
			m.authUsername = &username
			m.updateUserSpecificViews()
		}
	case profileState:
		m.profileView, cmd = m.profileView.Update(msg)
	case signupState:
		m.signUpView, cmd = m.signUpView.Update(msg)
	case feedState:
		m.feedView, cmd = m.feedView.Update(msg)
	case passwordResetState:
		m.passwordResetView, cmd = m.passwordResetView.Update(msg)
	case chatState:
		m.chatView, cmd = m.chatView.Update(msg)
	case publishPostState:
		m.publishPostView, cmd = m.publishPostView.Update(msg)
	}
	return m, cmd
}

func (m *mainView) updateUserSpecificViews() {
	m.profileView = profile.New(m.model, m.authUsername)
	m.feedView, _ = m.feedView.Update(util.UpdateUsername{Username: m.authUsername})
	m.passwordResetView = newPasswordResetView(&m.ctx, m.q, m.authUsername)
	m.chatView = chat.NewChatListView(m.authUsername, m.model)
	m.publishPostView = post.New(m.authUsername, m.model)
}

func (m mainView) View() string {
	sb := strings.Builder{}
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
	case chatState:
		sb.WriteString(m.chatView.View())
	case publishPostState:
		sb.WriteString(m.publishPostView.View())
	}
	if m.state != feedState {
		sb.WriteString(m.help.View()) // we can't render both the viewport  and the help component
		// this is a bit of an hack and can surely be achieved in a cleaner way
	}
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

func WithModel(model model.Model) viewOption {
	return func(m *mainView) {
		m.model = &model
	}
}

type loginView interface {
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
	GetAuthenticatedUsername() (string, error)
}
