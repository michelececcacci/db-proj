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
	dbname   = "postgres"
	user     = "postgres"
	password = "postgres"
	host     = "db"
	port     = "5432"
)

// just used to make sure our connection to the db is set, will be removed later
func testDbConnection(qq *queries.Queries, ctx context.Context) error {
	return qq.InsertUser(ctx, queries.InsertUserParams{})
}

func run() error {
	s := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", dbname, user, password, host, port)
	db, err := sql.Open("postgres", s)
	if err != nil {
		return err
	}
	qq := queries.New(db)
	err = testDbConnection(qq, context.Background())
	if err != nil {
		return err
	}
	p := tea.NewProgram(view.NewMainView(
		view.WithContext(context.Background()),
		view.WithQueries(qq),
	))
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
