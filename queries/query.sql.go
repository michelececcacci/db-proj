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

const getMemberId = `-- name: GetMemberId :one
SELECT IdMembro
FROM MEMBRO
WHERE DataEntrata = $1 AND Username = $2 AND IdChat = $3
`

type GetMemberIdParams struct {
	Dataentrata time.Time
	Username    string
	Idchat      int32
}

func (q *Queries) GetMemberId(ctx context.Context, arg GetMemberIdParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, getMemberId, arg.Dataentrata, arg.Username, arg.Idchat)
	var idmembro int32
	err := row.Scan(&idmembro)
	return idmembro, err
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

const getRandomAdmin = `-- name: GetRandomAdmin :one
SELECT IdMembro
FROM AMMINISTRATORE
ORDER BY random()
LIMIT 1
`

func (q *Queries) GetRandomAdmin(ctx context.Context) (int32, error) {
	row := q.db.QueryRowContext(ctx, getRandomAdmin)
	var idmembro int32
	err := row.Scan(&idmembro)
	return idmembro, err
}

const getRandomAdminInChat = `-- name: GetRandomAdminInChat :one
SELECT a.IdMembro
  FROM AMMINISTRATORE a JOIN MEMBRO m ON (a.IdMembro = m.IdMembro)
  WHERE m.IdChat = $1
  ORDER BY random()
  LIMIT 1
`

func (q *Queries) GetRandomAdminInChat(ctx context.Context, idchat int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, getRandomAdminInChat, idchat)
	var idmembro int32
	err := row.Scan(&idmembro)
	return idmembro, err
}

const getRandomChat = `-- name: GetRandomChat :one
SELECT IdChat
  FROM CHAT
  ORDER BY random()
  LIMIT 1
`

func (q *Queries) GetRandomChat(ctx context.Context) (int32, error) {
	row := q.db.QueryRowContext(ctx, getRandomChat)
	var idchat int32
	err := row.Scan(&idchat)
	return idchat, err
}

const getRandomMember = `-- name: GetRandomMember :one
SELECT IdMembro
FROM MEMBRO
ORDER BY random()
LIMIT 1
`

func (q *Queries) GetRandomMember(ctx context.Context) (int32, error) {
	row := q.db.QueryRowContext(ctx, getRandomMember)
	var idmembro int32
	err := row.Scan(&idmembro)
	return idmembro, err
}

const getRandomUser = `-- name: GetRandomUser :one

SELECT Username
FROM UTENTE
ORDER BY random()
LIMIT 1
`

// random accesses
func (q *Queries) GetRandomUser(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getRandomUser)
	var username string
	err := row.Scan(&username)
	return username, err
}

const insertAdmin = `-- name: InsertAdmin :exec
INSERT INTO AMMINISTRATORE (
  IdMembro
) VALUES ( $1 )
`

func (q *Queries) InsertAdmin(ctx context.Context, idmembro int32) error {
	_, err := q.db.ExecContext(ctx, insertAdmin, idmembro)
	return err
}

const insertChat = `-- name: InsertChat :one
INSERT INTO CHAT (Nome, Descrizione)
  VALUES ($1, $2)
  RETURNING IdChat
`

type InsertChatParams struct {
	Nome        string
	Descrizione string
}

func (q *Queries) InsertChat(ctx context.Context, arg InsertChatParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertChat, arg.Nome, arg.Descrizione)
	var idchat int32
	err := row.Scan(&idchat)
	return idchat, err
}

const insertFollower = `-- name: InsertFollower :exec
INSERT INTO SEGUIRE (UsernameSeguace, UsernameSeguito, DataInizio, DataFine)
  VALUES ($1, $2, $3, $4)
`

type InsertFollowerParams struct {
	Usernameseguace string
	Usernameseguito string
	Datainizio      time.Time
	Datafine        sql.NullTime
}

func (q *Queries) InsertFollower(ctx context.Context, arg InsertFollowerParams) error {
	_, err := q.db.ExecContext(ctx, insertFollower,
		arg.Usernameseguace,
		arg.Usernameseguito,
		arg.Datainizio,
		arg.Datafine,
	)
	return err
}

const insertMember = `-- name: InsertMember :one
INSERT INTO MEMBRO (
  DataEntrata, Username, IdChat, Amministratore
) VALUES ( $1, $2, $3, $4 )
RETURNING IdMembro
`

type InsertMemberParams struct {
	Dataentrata    time.Time
	Username       string
	Idchat         int32
	Amministratore sql.NullInt32
}

func (q *Queries) InsertMember(ctx context.Context, arg InsertMemberParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertMember,
		arg.Dataentrata,
		arg.Username,
		arg.Idchat,
		arg.Amministratore,
	)
	var idmembro int32
	err := row.Scan(&idmembro)
	return idmembro, err
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

const insertRegion = `-- name: InsertRegion :exec
INSERT INTO REGIONE (Nome, Superregione) 
  VALUES ($1, $2)
`

type InsertRegionParams struct {
	Nome         string
	Superregione sql.NullInt32
}

func (q *Queries) InsertRegion(ctx context.Context, arg InsertRegionParams) error {
	_, err := q.db.ExecContext(ctx, insertRegion, arg.Nome, arg.Superregione)
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

const isValidAdmin = `-- name: IsValidAdmin :one
SELECT COUNT(m.IdMembro)
FROM AMMINISTRATORE a JOIN MEMBRO m ON (a.IdMembro = m.IdMembro)
WHERE a.IdMembro = $1 AND m.IdChat = $2
`

type IsValidAdminParams struct {
	Idmembro int32
	Idchat   int32
}

func (q *Queries) IsValidAdmin(ctx context.Context, arg IsValidAdminParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, isValidAdmin, arg.Idmembro, arg.Idchat)
	var count int64
	err := row.Scan(&count)
	return count, err
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
