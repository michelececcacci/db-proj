-- name: TestQuery :one
SELECT * FROM CHAT;

-- name: InsertUser :exec
INSERT INTO UTENTE (Username, DataDiNascita, Nome, Cognome, Domicilio) 
    VALUES             ($1,    $2,        $3,   $4,      $5);

-- name: InsertPassword :exec
INSERT INTO STORICO_PASSWORD (Username, Password, DataInserimento)
    VALUES ($1, $2, $3);

-- name: InsertFollower :exec
INSERT INTO SEGUIRE (UsernameSeguace, UsernameSeguito, DataInizio, DataFine)
  VALUES ($1, $2, $3, $4);

-- name: InsertRegion :exec
INSERT INTO REGIONE (Nome, Superregione) 
  VALUES ($1, $2);

-- name: InsertChat :one
INSERT INTO CHAT (Nome, Descrizione)
  VALUES ($1, $2)
  RETURNING IdChat;

  -- name: InsertMember :one
INSERT INTO MEMBRO (
  DataEntrata, Username, IdChat, Amministratore
) VALUES ( $1, $2, $3, $4 )
RETURNING IdMembro;


-- name: InsertAdmin :exec
INSERT INTO AMMINISTRATORE (
  IdMembro
) VALUES ( $1 );

-- name: GetFollowers :many
SELECT usernameseguace FROM SEGUIRE WHERE usernameseguito = $1 AND DataFine IS NULL;

-- name: GetFollowing :many
SELECT usernameseguito FROM SEGUIRE WHERE usernameseguace = $1 AND DataFine IS NULL;

-- name: GetMemberId :one
SELECT IdMembro
FROM MEMBRO
WHERE DataEntrata = $1 AND Username = $2 AND IdChat = $3;

-- name: IsValidAdmin :one
SELECT COUNT(m.IdMembro)
FROM AMMINISTRATORE a JOIN MEMBRO m ON (a.IdMembro = m.IdMembro)
WHERE a.IdMembro = $1 AND m.IdChat = $2;

-- name: Authenticate :one
SELECT COUNT(SUBQUERY.DataInserimento)
FROM (
   SELECT MIN(SP.DataInserimento) AS DataInserimento
   FROM STORICO_PASSWORD SP 
   WHERE SP.Username = $1 AND SP.Password = $2
) AS SUBQUERY;


-- name: GetPastPasswords :many
SELECT Password FROM  STORICO_PASSWORD WHERE username = $1;

-- name: GetLocations :many
SELECT * FROM Regione;
-- random accesses

-- name: GetRandomUser :one
SELECT Username
FROM UTENTE
ORDER BY random()
LIMIT 1;

-- name: GetRandomChat :one 
SELECT IdChat
  FROM CHAT
  ORDER BY random()
  LIMIT 1;

-- name: GetRandomMember :one
SELECT IdMembro
FROM MEMBRO
ORDER BY random()
LIMIT 1;

-- name: GetRandomAdmin :one
SELECT IdMembro
FROM AMMINISTRATORE
ORDER BY random()
LIMIT 1;

-- name: GetRandomAdminInChat :one 
SELECT a.IdMembro
  FROM AMMINISTRATORE a JOIN MEMBRO m ON (a.IdMembro = m.IdMembro)
  WHERE m.IdChat = $1
  ORDER BY random()
  LIMIT 1;
