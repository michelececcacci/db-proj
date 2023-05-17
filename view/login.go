package view

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/michelececcacci/db-proj/util"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	
	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type loginView struct {
	currentElement int
	inputs         []textinput.Model
}

func (l loginView) View() string {
	var b strings.Builder
	for _, input := range l.inputs {
		b.WriteString(input.View())
		b.WriteString("\n")
	}
	return b.String()
}

func (l loginView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return l, tea.Quit
		case tea.KeyDown:
			l.currentElement = util.Min(len(l.inputs)-1, l.currentElement+1)
		case tea.KeyUp:
			l.currentElement = util.Max(0, l.currentElement-1)
		}
		l.inputs[l.currentElement], cmd = l.inputs[l.currentElement].Update(msg)
		for i := range(l.inputs) {
			if i == l.currentElement {
				focusButton(&l.inputs[i])
			} else {
				blurButton(&l.inputs[i])
			}
		}
	}
	return l, cmd
}

func (l loginView) Init() tea.Cmd {
	return textinput.Blink
}

func newLoginView() loginView {
	inputs := make([]textinput.Model, 2)
	inputs[0] = newInput("username", 20)
	inputs[1] = newInput("password", 20)
	focusButton(&inputs[0])
	m := loginView{
		inputs: inputs,
	}
	return m
}

func newInput(placeholder string, maxLenght int) textinput.Model {
	t := textinput.New()
	t.Placeholder = placeholder
	t.CharLimit = maxLenght
	return t
}

func focusButton(t * textinput.Model) {
	t.Focus()
	t.PromptStyle = focusedStyle
	t.TextStyle = focusedStyle
}

func blurButton(t *textinput.Model) {
	t.Blur()
	t.PromptStyle = blurredStyle
	t.TextStyle = blurredStyle
}