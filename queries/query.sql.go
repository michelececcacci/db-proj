// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package queries

import (
	"context"
	"database/sql"
	"time"
)

const authenticate = `-- name: Authenticate :one
SELECT COUNT(SUBQUERY.DataInserimento)
FROM (
   SELECT MIN(SP.DataInserimento) AS DataInserimento
   FROM STORICO_PASSWORD SP 
   WHERE SP.Username = $1 AND SP.Password = $2
) AS SUBQUERY
`

type AuthenticateParams struct {
	Username string
	Password string
}

func (q *Queries) Authenticate(ctx context.Context, arg AuthenticateParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, authenticate, arg.Username, arg.Password)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getFollowers = `-- name: GetFollowers :many
SELECT usernameseguace FROM SEGUIRE WHERE usernameseguito = $1 AND DataFine IS NULL
`

func (q *Queries) GetFollowers(ctx context.Context, usernameseguito string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getFollowers, usernameseguito)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var usernameseguace string
		if err := rows.Scan(&usernameseguace); err != nil {
			return nil, err
		}
		items = append(items, usernameseguace)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowing = `-- name: GetFollowing :many
SELECT usernameseguito FROM SEGUIRE WHERE usernameseguace = $1 AND DataFine IS NULL
`

func (q *Queries) GetFollowing(ctx context.Context, usernameseguace string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getFollowing, usernameseguace)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var usernameseguito string
		if err := rows.Scan(&usernameseguito); err != nil {
			return nil, err
		}
		items = append(items, usernameseguito)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLocations = `-- name: GetLocations :many
SELECT idregione, nome, superregione FROM Regione
`

func (q *Queries) GetLocations(ctx context.Context) ([]Regione, error) {
	rows, err := q.db.QueryContext(ctx, getLocations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Regione
	for rows.Next() {
		var i Regione
		if err := rows.Scan(&i.Idregione, &i.Nome, &i.Superregione); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPastPasswords = `-- name: GetPastPasswords :many
SELECT Password FROM  STORICO_PASSWORD WHERE username = $1
`

func (q *Queries) GetPastPasswords(ctx context.Context, username string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getPastPasswords, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var password string
		if err := rows.Scan(&password); err != nil {
			return nil, err
		}
		items = append(items, password)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertPassword = `-- name: InsertPassword :exec
INSERT INTO STORICO_PASSWORD (Username, Password, DataInserimento)
    VALUES ($1, $2, $3)
`

type InsertPasswordParams struct {
	Username        string
	Password        string
	Datainserimento time.Time
}

func (q *Queries) InsertPassword(ctx context.Context, arg InsertPasswordParams) error {
	_, err := q.db.ExecContext(ctx, insertPassword, arg.Username, arg.Password, arg.Datainserimento)
	return err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO UTENTE (Username, DataDiNascita, Nome, Cognome, Domicilio) 
    VALUES             ($1,    $2,        $3,   $4,      $5)
`

type InsertUserParams struct {
	Username      string
	Datadinascita sql.NullTime
	Nome          sql.NullString
	Cognome       sql.NullString
	Domicilio     sql.NullInt32
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.db.ExecContext(ctx, insertUser,
		arg.Username,
		arg.Datadinascita,
		arg.Nome,
		arg.Cognome,
		arg.Domicilio,
	)
	return err
}

const testQuery = `-- name: TestQuery :one
SELECT nome, descrizione, idchat FROM CHAT
`

func (q *Queries) TestQuery(ctx context.Context) (Chat, error) {
	row := q.db.QueryRowContext(ctx, testQuery)
	var i Chat
	err := row.Scan(&i.Nome, &i.Descrizione, &i.Idchat)
	return i, err
}
