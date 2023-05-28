-- name: TestQuery :one
SELECT * FROM CHAT;

-- name: InsertUser :exec
INSERT INTO UTENTE (Username, DataDiNascita, Nome, Cognome, Domicilio) 
    VALUES             ($1,    $2,        $3,   $4,      $5);

-- name: InsertPassword :exec
INSERT INTO STORICO_PASSWORD (Username, Password, DataInserimento)
    VALUES ($1, $2, $3);

-- name: GetFollowers :many
SELECT usernameseguace FROM SEGUIRE WHERE usernameseguito = $1 AND DataFine IS NULL;

-- name: GetFollowing :many
SELECT usernameseguito FROM SEGUIRE WHERE usernameseguace = $1 AND DataFine IS NULL;

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