-- name: TestQuery :one
SELECT * FROM CHAT;

-- name: InsertUser :exec
INSERT INTO UTENTE (Username, DataDiNascita, Nome, Cognome, Domicilio) 
    VALUES             ($1,    $2,        $3,   $4,      $5);

-- name: InsertPassword :exec
INSERT INTO STORICO_PASSWORD (Username, Password, DataInserimento)
    VALUES ($1, $2, $3);

-- name: GetFollowers :many
SELECT usernameseguace FROM SEGUIRE WHERE usernameseguito = $1;

-- name: GetFollwing :many
SELECT usernameseguito FROM SEGUIRE WHERE usernameseguace = $1;
