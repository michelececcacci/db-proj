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

-- name: InsertExit :exec
INSERT INTO USCITA (
IdMembro, DataUscita, Motivazione, IdMembroResponsabile
) VALUES ($1, $2, $3, $4); 

-- name: GetFollowers :many
SELECT usernameseguace FROM SEGUIRE WHERE usernameseguito = $1 AND DataFine IS NULL;

-- name: GetFollowing :many
SELECT usernameseguito FROM SEGUIRE WHERE usernameseguace = $1 AND DataFine IS NULL;

-- name: GetMemberId :one
SELECT IdMembro
FROM MEMBRO
WHERE DataEntrata = $1 AND Username = $2 AND IdChat = $3;

-- name: GetCurrentMember :one
SELECT IdMembro, DataEntrata
FROM MEMBRO
WHERE Username = $1 AND IdChat = $2
ORDER BY DataEntrata ASC
LIMIT 1;

-- name: GetDataOfMember :one 
SELECT username, idchat, DataEntrata
FROM MEMBRO
WHERE IdMembro = $1;

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


-- name: CheckIfUserStillInChat :one
SELECT COUNT(*)
FROM MEMBRO M FULL OUTER JOIN USCITA U ON (M.IDMEMBRO = U.IDMEMBRO)
WHERE M.USERNAME = $1
	AND M.IDCHAT = $2
	AND U.IDMEMBRO IS NULL;

-- name: CheckIfMemberStillInChat :one 
SELECT M.DataEntrata
FROM MEMBRO M FULL OUTER JOIN USCITA U ON (M.IDMEMBRO = U.IDMEMBRO)
WHERE M.IDMEMBRO = $1
	AND U.IDMEMBRO IS NULL;


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

-- name: GetAllPossibleMembers :many 
SELECT u.username AS username, c.IdChat AS idchat
FROM UTENTE u FULL OUTER JOIN CHAT c ON (True)
ORDER BY random();

-- name: GetAllPossibleFollowings :many 
SELECT u1.username AS U1, u2.username AS U2
FROM UTENTE u1 FULL OUTER JOIN UTENTE u2 ON (True)
ORDER BY random();

-- name: GetRandomMemberInChat :one
SELECT m.IdMembro, m.DataEntrata
FROM MEMBRO m full outer join USCITA u ON (m.IdMembro = u.IdMembro)
WHERE u.IdMembro IS NULL AND m.IdChat = $1;

-- name: GetRandomMemberFromAnyChat :many 
SELECT m.IdMembro, m.IdChat
FROM MEMBRO m 
ORDER BY random();

-- name: GetLocationRec :many
WITH recursive getSuperregions(idregione, superregione)
AS(
	(
		select idregione, superregione
		from regione
	) union all (
		select g.idregione, a.superregione
		from regione g, getSuperregions a
		where g.superregione = a.idregione
	)
)

select nome
from getSuperregions g join regione r on (g.superregione = r.idregione)
where g.idregione = $1;

-- name: GetFullFeed :many
SELECT titolo, autore, TimestampPubblicazione, Testo FROM SEGUIRE 
JOIN CONTENUTO ON CONTENUTO.Autore = SEGUIRE.usernameSeguito  
WHERE SEGUIRE.usernameseguace = ($1) AND datafine IS NULL AND IdContenutoPadre IS NULL;


-- name: GetSpecificPost :many
SELECT * FROM CONTENUTO WHERE Autore = $1 AND IdContenuto = $2;

-- name: InsertMessage :exec
INSERT INTO MESSAGGIO (Testo, TimestampInvio, Mittente) VALUES ($1, $2, $3);

-- name: GetChatIds :many
SELECT IdMembro, IdChat FROM MEMBRO WHERE Username = $1;

-- name: GetChatInfos :one
SELECT Nome, Descrizione FROM CHAT WHERE idChat = $1;

-- name: GetChatMessages :many
SELECT testo, TimestampInvio, username FROM MESSAGGIO JOIN MEMBRO ON MEMBRO.IdMembro = MESSAGGIO.Mittente WHERE MEMBRO.idChat = $1;

-- name: GetChatMembers :many
SELECT IdMembro FROM MEMBRO WHERE MEMBRO.idchat = $1;