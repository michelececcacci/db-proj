package signup

import (
	"database/sql"

	"github.com/charmbracelet/bubbles/list"
	"github.com/michelececcacci/db-proj/queries"
)

type location struct {
	id     int32
	name   string
	parent sql.NullInt32
}

func (l location) Title() string {
	return l.name
}

func (l location) FilterValue() string {
	return l.Title()
}

func (l location) Description() string {
	return ""
}

func toLocations(r []queries.Regione) []list.Item {
	var l []list.Item
	for _, region := range r {
		l = append(l, location{
			id:     region.Idregione,
			parent: region.Superregione,
			name:   region.Nome,
		})
	}
	return l
}
