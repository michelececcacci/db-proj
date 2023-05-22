package components

import "github.com/charmbracelet/bubbles/textinput"

func NewTextInput(placeHolder string, charLimit int) textinput.Model {
	t := textinput.New()
	t.Placeholder = placeHolder
	t.CharLimit = charLimit
	return t
}
