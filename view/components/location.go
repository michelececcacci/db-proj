package components

import (
	"context"
	"database/sql"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/michelececcacci/db-proj/queries"
)

type location struct {
	id          int32
	name        string
	parent      sql.NullInt32
	description string
}

func (l location) Title() string {
	return l.name
}

func (l location) FilterValue() string {
	return l.Title()
}

func (l location) Description() string {
	return l.description
}

func toLocations(r []queries.Regione, q *queries.Queries, ctx *context.Context) []list.Item {
	var l []list.Item
	for _, region := range r {
		parents, _ := q.GetLocationRec(*ctx, region.Idregione)
		l = append(l, location{
			id:          region.Idregione,
			parent:      region.Superregione,
			name:        region.Nome,
			description: strings.Join(parents, ","),
		})
	}
	return l
}
