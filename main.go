package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/michelececcacci/db-proj/queries"
)

func run() error {
	db, err := sql.Open("postgres", "user=postgres dbname=logica sslmode=verify-full")
	if err != nil {
		return err
	}
	qq := queries.New(db)
	r, err := qq.Something(context.Background())
	if err != nil {
		return err
	}
	log.Println(r)
	// p := tea.NewProgram(initialModel())
	// _, err := p.Run()
	// if err != nil {
	// 	panic(err)
	// }
	return nil
}

func main() {
	log.Println("hi")
	// if err :=run(); err != nil {
	// 	log.Fatal(err)
	// }
}
