package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/lib/pq"
	"github.com/michelececcacci/db-proj/queries"
	"github.com/michelececcacci/db-proj/view"
)

const (
	dbname="db"
	user="postgres"
	password="postgres"
	host="db"
)

func run() error {
	s := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=disable",dbname, user, password, host)
	db, err := sql.Open("postgres", s)
	if err != nil {
		return err
	}
	qq := queries.New(db)
	err = qq.InsertUser(context.Background(), "hi")
	if err != nil {
		return err
	}
	p := tea.NewProgram(view.InitialModel())
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
