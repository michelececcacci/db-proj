-- *********************************************
-- * SQL PostgreSQL generation                 
-- *--------------------------------------------
-- * DB-MAIN version: 11.0.2              
-- * Generator date: Sep 20 2021              
-- * Generation date: Sat May 20 17:17:32 2023 
-- * LUN file: /home/luca/db-proj/Social Network.lun 
-- * Schema: Logica/1-1 
-- ********************************************* 


-- Database Section
-- ________________ 



-- Tables Section
-- _____________ 

create table AMMINISTRATORE (
     IdMembro numeric(1) not null,
     constraint FKAMMINISTRATORE_ID primary key (IdMembro));

create table CHAT (
     Nome varchar(100) not null,
     Descrizione varchar(5000) not null,
     IdChat numeric(1) not null,
     constraint ID_CHAT_ID primary key (IdChat));

create table CONTENUTO (
     Autore varchar(20) not null,
     Testo varchar(1000) not null,
     IdContenuto numeric(1) not null,
     TimestampPubblicazione date not null,
     Titolo varchar(30),
     IdRegione numeric(1),
     UsernamePadre varchar(20),
     IdContenutoPadre numeric(1),
     constraint ID_CONTENUTO primary key (Autore, IdContenuto));

create table MEMBRO (
     DataEntrata date not null,
     IdMembro numeric(1) not null,
     Username varchar(20) not null,
     IdChat numeric(1) not null,
     Amministratore numeric(1),
     constraint ID_MEMBRO primary key (IdMembro),
     constraint SID_MEMBRO unique (IdChat, Username, DataEntrata));

create table MESSAGGIO (
     IdMessaggio numeric(1) not null,
     Testo varchar(1000) not null,
     TimestampInvio date not null,
     Mittente numeric(1) not null,
     Citazione numeric(1),
     AutoreContenuto varchar(20),
     IdContenuto numeric(1),
     AmministratoreEliminatore numeric(1),
     constraint ID_MESSAGGIO primary key (IdMessaggio));

create table REAZIONE (
     AutoreContenuto varchar(20) not null,
     IdContenuto numeric(1) not null,
     Username varchar(20) not null,
     LikeDislike numeric(1) not null,
     Timestamp char(1) not null,
     constraint ID_REAZIONE primary key (Username, AutoreContenuto, IdContenuto));

create table REGIONE (
     IdRegione numeric(1) not null,
     Nome varchar(50) not null,
     Superregione numeric(1),
     constraint ID_REGIONE primary key (IdRegione));

create table SEGUIRE (
     UsernameSeguace varchar(20) not null,
     UsernameSeguito varchar(20) not null,
     DataInizio date not null,
     DataFine date,
     constraint ID_SEGUIRE primary key (UsernameSeguace, UsernameSeguito, DataInizio));

create table STORICO_ACCESSO (
     Username varchar(20) not null,
     TimestampLogin date not null,
     TimestampLogout date not null,
     constraint ID_STORICO_ACCESSO primary key (Username, TimestampLogin));

create table STORICO_PASSWORD (
     Username varchar(20) not null,
     DataInserimento date not null,
     Password varchar(20) not null,
     constraint SID_STORICO_PASSWORD unique (Username, DataInserimento),
     constraint ID_STORICO_PASSWORD primary key (Username, Password));

create table USCITA (
     IdMembro numeric(1) not null,
     DataUscita date not null,
     Motivazione varchar(500),
     IdMembroResponsabile numeric(1),
     constraint FKUSCITA_ID primary key (IdMembro));

create table UTENTE (
     Username varchar(20) not null,
     DataDiNascita date not null,
     Nome varchar(20) not null,
     Cognome varchar(20) not null,
     Domicilio numeric(1),
     constraint ID_UTENTE_ID primary key (Username));


-- Constraints Section
-- ___________________ 

alter table AMMINISTRATORE add constraint FKAMMINISTRATORE_FK
     foreign key (IdMembro)
     references MEMBRO;

--Not implemented
--alter table CHAT add constraint ID_CHAT_CHK
--     check(exists(select * from MEMBRO
--                  where MEMBRO.IdChat = IdChat)); 

alter table CONTENUTO add constraint FKPUBBLICAZIONE
     foreign key (Autore)
     references UTENTE;

alter table CONTENUTO add constraint FKLUOGO_PUBBLICAZIONE_FK
     foreign key (IdRegione)
     references REGIONE;

alter table CONTENUTO add constraint FKRISPOSTA_FK
     foreign key (UsernamePadre, IdContenutoPadre)
     references CONTENUTO;

alter table CONTENUTO add constraint FKRISPOSTA_CHK
     check((UsernamePadre is not null and IdContenutoPadre is not null)
           or (UsernamePadre is null and IdContenutoPadre is null)); 

alter table MEMBRO add constraint FKIMPERSONAZIONE_FK
     foreign key (Username)
     references UTENTE;

alter table MEMBRO add constraint FKAPPARTENENZA
     foreign key (IdChat)
     references CHAT;

alter table MEMBRO add constraint FKAGGIUNTO_FK
     foreign key (Amministratore)
     references AMMINISTRATORE;

alter table MESSAGGIO add constraint FKMITTENTE_FK
     foreign key (Mittente)
     references MEMBRO;

alter table MESSAGGIO add constraint FKCITAZIONE_FK
     foreign key (Citazione)
     references MESSAGGIO;

alter table MESSAGGIO add constraint FKRIFERIMENTO_FK
     foreign key (AutoreContenuto, IdContenuto)
     references CONTENUTO;

alter table MESSAGGIO add constraint FKRIFERIMENTO_CHK
     check((AutoreContenuto is not null and IdContenuto is not null)
           or (AutoreContenuto is null and IdContenuto is null)); 

alter table MESSAGGIO add constraint FKELIMINAZIONE_FK
     foreign key (AmministratoreEliminatore)
     references AMMINISTRATORE;

alter table REAZIONE add constraint FKREA_UTE
     foreign key (Username)
     references UTENTE;

alter table REAZIONE add constraint FKREA_CON_FK
     foreign key (AutoreContenuto, IdContenuto)
     references CONTENUTO;

alter table REGIONE add constraint FKPARTE_DI_FK
     foreign key (Superregione)
     references REGIONE;

alter table SEGUIRE add constraint FKSEGUITO_FK
     foreign key (UsernameSeguito)
     references UTENTE;

alter table SEGUIRE add constraint FKSEGUE_FK
     foreign key (UsernameSeguace)
     references UTENTE;

alter table STORICO_ACCESSO add constraint FKACCESSO
     foreign key (Username)
     references UTENTE;

alter table STORICO_PASSWORD add constraint FKSCELTA
     foreign key (Username)
     references UTENTE;

alter table USCITA add constraint FKUSCITA_FK
     foreign key (IdMembro)
     references MEMBRO;

alter table USCITA add constraint FKRESPONSABILE_FK
     foreign key (IdMembroResponsabile)
     references AMMINISTRATORE;

--Not implemented
--alter table UTENTE add constraint ID_UTENTE_CHK
--     check(exists(select * from STORICO_PASSWORD
--                  where STORICO_PASSWORD.Username = Username)); 

alter table UTENTE add constraint FKDOMICILIO_FK
     foreign key (Domicilio)
     references REGIONE;


-- Index Section
-- _____________ 

create index FKLUOGO_PUBBLICAZIONE_IND
     on CONTENUTO (IdRegione);

create index FKRISPOSTA_IND
     on CONTENUTO (UsernamePadre, IdContenutoPadre);

create index FKIMPERSONAZIONE_IND
     on MEMBRO (Username);

create index FKAGGIUNTO_IND
     on MEMBRO (Amministratore);

create index FKMITTENTE_IND
     on MESSAGGIO (Mittente);

create index FKCITAZIONE_IND
     on MESSAGGIO (Citazione);

create index FKRIFERIMENTO_IND
     on MESSAGGIO (AutoreContenuto, IdContenuto);

create index FKELIMINAZIONE_IND
     on MESSAGGIO (AmministratoreEliminatore);

create index FKREA_CON_IND
     on REAZIONE (AutoreContenuto, IdContenuto);

create index FKPARTE_DI_IND
     on REGIONE (Superregione);

create index FKSEGUITO_IND
     on SEGUIRE (UsernameSeguito);

create index FKSEGUE_IND
     on SEGUIRE (UsernameSeguace);

create index FKRESPONSABILE_IND
     on USCITA (IdMembroResponsabile);

create index FKDOMICILIO_IND
     on UTENTE (Domicilio);

