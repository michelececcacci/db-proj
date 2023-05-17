package main

import (
	"context"
	"database/sql"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/lib/pq"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/view"
)

func run() error {
	db, err := sql.Open("postgres", "user=postgres dbname=logica sslmode=verify-full")
	if err != nil {
		return err
	}
	qq := queries.New(db)
	_, err = qq.Something(context.Background())
	if err != nil {
		return err
	}
	p := tea.NewProgram(view.InitialModel())
	_, err = p.Run()
	if err != nil {
		panic(err)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
