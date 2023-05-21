package components

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/michelececcacci/db-proj/util"
)

// Treated as constant
var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)

type MultipleInputsView struct {
	currentElement int
	inputs         []textinput.Model // TODO this might get changed to a pointer
}

func (iv MultipleInputsView) Init() tea.Cmd {
	return textinput.Blink
}

func (iv MultipleInputsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return iv, tea.Quit
		case tea.KeyDown, tea.KeyTab:
			iv.currentElement = util.Min(len(iv.inputs)-1, iv.currentElement+1)
		case tea.KeyUp, tea.KeyShiftTab:
			iv.currentElement = util.Max(0, iv.currentElement-1)
		default:
			iv.inputs[iv.currentElement], cmd = iv.inputs[iv.currentElement].Update(msg)
		}
		for idx := range iv.inputs {
			if idx == iv.currentElement {
				focusButton(&iv.inputs[idx])
			} else {
				blurButton(&iv.inputs[idx])
			}
		}
	}
	return iv, cmd
}

func NewInput(placeholder string, maxLenght int) textinput.Model {
	t := textinput.New()
	t.Placeholder = placeholder
	t.CharLimit = maxLenght
	return t
}

func focusButton(t *textinput.Model) {
	t.Focus()
	t.PromptStyle = focusedStyle
	t.TextStyle = focusedStyle
}

func blurButton(t *textinput.Model) {
	t.Blur()
	t.PromptStyle = blurredStyle
	t.TextStyle = blurredStyle
}

func (iv MultipleInputsView) View() string {
	var b strings.Builder
	b.WriteString("Press Enter to submit\n")
	for _, input := range iv.inputs {
		b.WriteString(input.View())
		b.WriteString("\n")
	}
	return b.String()
}

func NewMultipleInputsView(fields []textinput.Model) MultipleInputsView {
	if len(fields) == 0 {
		panic("Inputs view fields array can't be empty")
	}
	iv := MultipleInputsView{}
	iv.inputs = fields
	iv.currentElement = 0 // rendundant but surely more explicit
	focusButton(&iv.inputs[0])
	return iv
}

func (iv MultipleInputsView) GetInputValueByIndex(i int) string {
	return iv.inputs[i].Value()
}
