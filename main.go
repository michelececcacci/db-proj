package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/lib/pq"
	"github.com/michelececcacci/db-proj/model"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/view"
)

const (
	dbname   = "postgres"
	user     = "postgres"
	password = "postgres"
)

var (
	host = os.Getenv("HOSTNAME")
	port = os.Getenv("PORTNAME")
)

func run() error {
	s := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", dbname, user, password, host, port)
	db, err := sql.Open("postgres", s)
	if err != nil {
		return err
	}
	qq := queries.New(db)
	m := model.New(db, context.Background())
	p := tea.NewProgram(
		view.NewMainView(
			view.WithContext(context.Background()),
			view.WithQueries(qq),
			view.WithModel(m),
		),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	_, err = p.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
