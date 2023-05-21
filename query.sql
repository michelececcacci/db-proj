-- name: TestQuery :one
SELECT * FROM CHAT;

-- name: InsertUser :exec
INSERT INTO UTENTE (Username, DataDiNascita, Nome, Cognome, Domicilio) 
VALUES             ($1,       $2,            $3,   $4,      $5);