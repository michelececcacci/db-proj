package components

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/michelececcacci/db-proj/queries"
)

type SelectLocation struct {
	SelectedLocation sql.NullInt32
	locations        list.Model
}

func (s SelectLocation) Init() tea.Cmd {
	return nil
}

func (s SelectLocation) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			s.SelectedLocation.Valid = true
			selected := s.locations.SelectedItem().FilterValue()
			id, _ := strconv.ParseInt(selected, 10, 32)
			s.SelectedLocation.Int32 = int32(id)
		default:
			s.locations , cmd = s.locations.Update(msg)
		}
	}
	return s, cmd
}

func (s SelectLocation) View() string {
	return s.locations.View()
}

func NewSelectLocation(q *queries.Queries, ctx *context.Context) SelectLocation {
	result, err := q.GetLocations(*ctx)
	if err != nil {
		panic(err)
	}
	locations := toLocations(result, q, ctx)
	s := SelectLocation{
		locations: list.New(locations, list.NewDefaultDelegate(), 25, 25),
	}
	s.locations.Title = "Select location"
	return s
}

