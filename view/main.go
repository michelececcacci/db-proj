package view

// Subviews implement tea.Model and the main view is
// responsible for updating them and selecting the displayed one
// according to the state changes.

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
)

type viewOption func(*mainView)

type mainView struct {
	ctx          context.Context
	sqlcQueries  *queries.Queries // this needs to be renamed
	loginView    tea.Model
	profileView  tea.Model
	registerView tea.Model
	feedView     tea.Model
}

func NewMainView(options ...viewOption) mainView {
	m := mainView{
		loginView:    newLoginView(),
		profileView:  newProfileView(),
		registerView: newRegisterView(),
		feedView:     NewFeedView(),
	}
	for _, opt := range options {
		opt(&m)
	}
	return m
}

func (m mainView) Init() tea.Cmd {
	return nil
}

func (m mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// return m.loginView.Update(msg)
	// return m.profileView.Update(msg)
	// return m.registerView.Update(msg)
	return m.feedView.Update(msg)
}

func (m mainView) View() string {
	// return m.loginView.View()
	// return m.profileView.View()
	// return m.registerView.View()
	return m.feedView.View()
}

func WithContext(ctx context.Context) viewOption {
	return func(m *mainView) {
		m.ctx = ctx
	}
}

func WithQueries(q *queries.Queries) viewOption {
	return func(m *mainView) {
		m.sqlcQueries = q
	}
}
