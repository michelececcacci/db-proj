package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/michelececcacci/db-proj/util"
)

var errorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#fa1111"))

type errorView struct {
	message string
	err     error
}

func (e errorView) Init() tea.Cmd {
	return nil
}

func (e errorView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case util.OptionalError:
		e.message = msg.Message
		e.err = msg.Err
	}
	return e, cmd
}

// Returns a new errorView: a view that displays an error
// (if there is one) or otherwise displays a message string
func NewErrorView() errorView {
	return errorView{}
}

func (e errorView) View() string {
	if e.err != nil {
		return errorStyle.Render(e.err.Error() + "\n")
	}
	return e.message + "\n"
}
