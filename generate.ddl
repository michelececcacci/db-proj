-- *********************************************
-- * SQL PostgreSQL generation                 
-- *--------------------------------------------
-- * DB-MAIN version: 11.0.2              
-- * Generator date: Sep 20 2021              
-- * Generation date: Tue May 16 18:26:38 2023 
-- * LUN file: /home/michele/uni/db/db-proj/Social Network.lun 
-- * Schema: Logica/1-1 
-- ********************************************* 


-- Database Section
-- ________________ 

create database Logica;


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
     Username varchar(20) not null,
     Testo varchar(1000) not null,
     IdContenuto numeric(1) not null,
     TimestampPubblicazione timestamp not null,
     IdRegione numeric(1),
     A_Username varchar(20),
     A_IdContenuto numeric(1),
     constraint ID_CONTENUTO primary key (Username, IdContenuto));

create table MEMBRO (
     DataEntrata timestamp not null,
     IdMembro numeric(1) not null,
     Username varchar(20) not null,
     IdChat numeric(1) not null,
     Amministratore numeric(1),
     constraint ID_MEMBRO primary key (IdMembro),
     constraint SID_MEMBRO unique (IdChat, Username, DataEntrata));

create table MESSAGGIO (
     IdMessaggio numeric(1) not null,
     Testo varchar(1000) not null,
     TimestampInvio timestamp not null,
     IdMembro numeric(1) not null,
     Citazione numeric(1),
     Username varchar(20),
     IdContenuto numeric(1),
     AmministratoreEliminatore numeric(1),
     constraint ID_MESSAGGIO primary key (IdMessaggio));

create table REAZIONE (
     R_C_Username varchar(20) not null,
     R_C_IdContenuto numeric(1) not null,
     Username varchar(20) not null,
     LikeDislike numeric(1) not null,
     Timestamp char(1) not null,
     constraint ID_REAZIONE primary key (Username, R_C_Username, R_C_IdContenuto));

create table REGIONE (
     IdRegione numeric(1) not null,
     Nome varchar(50) not null,
     Superregione numeric(1),
     constraint ID_REGIONE primary key (IdRegione));

create table SEGUIRE (
     SEG_Username varchar(20) not null,
     Username varchar(20) not null,
     DataInizio timestamp not null,
     DataFine timestamp,
     constraint ID_SEGUIRE primary key (SEG_Username, Username, DataInizio));

create table STORICO_ACCESSO (
     Username varchar(20) not null,
     TimestampLogin timestamp not null,
     TimestampLogout timestamp not null,
     constraint ID_STORICO_ACCESSO primary key (Username, TimestampLogin));

create table STORICO_PASSWORD (
     Username varchar(20) not null,
     DataInserimento timestamp not null,
     Password varchar(20) not null,
     constraint SID_STORICO_PASSWORD unique (Username, DataInserimento),
     constraint ID_STORICO_PASSWORD primary key (Username, Password));

create table USCITA (
     IdMembro numeric(1) not null,
     DataUscita timestamp not null,
     Motivazione varchar(500),
     RES_IdMembro numeric(1),
     constraint FKUSCITA_ID primary key (IdMembro));

create table UTENTE (
     Username varchar(20) not null,
     Gen_DataDiNascita date not null,
     Gen_Nome varchar(20) not null,
     Gen_Cognome varchar(20) not null,
     IdRegione numeric(1),
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
     foreign key (Username)
     references UTENTE;

alter table CONTENUTO add constraint FKLUOGO_PUBBLICAZIONE_FK
     foreign key (IdRegione)
     references REGIONE;

alter table CONTENUTO add constraint FKRISPOSTA_FK
     foreign key (A_Username, A_IdContenuto)
     references CONTENUTO;

alter table CONTENUTO add constraint FKRISPOSTA_CHK
     check((A_Username is not null and A_IdContenuto is not null)
           or (A_Username is null and A_IdContenuto is null)); 

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
     foreign key (IdMembro)
     references MEMBRO;

alter table MESSAGGIO add constraint FKCITAZIONE_FK
     foreign key (Citazione)
     references MESSAGGIO;

alter table MESSAGGIO add constraint FKRIFERIMENTO_FK
     foreign key (Username, IdContenuto)
     references CONTENUTO;

alter table MESSAGGIO add constraint FKRIFERIMENTO_CHK
     check((Username is not null and IdContenuto is not null)
           or (Username is null and IdContenuto is null)); 

alter table MESSAGGIO add constraint FKELIMINAZIONE_FK
     foreign key (AmministratoreEliminatore)
     references AMMINISTRATORE;

alter table REAZIONE add constraint FKREA_UTE
     foreign key (Username)
     references UTENTE;

alter table REAZIONE add constraint FKREA_CON_FK
     foreign key (R_C_Username, R_C_IdContenuto)
     references CONTENUTO;

alter table REGIONE add constraint FKPARTE_DI_FK
     foreign key (Superregione)
     references REGIONE;

alter table SEGUIRE add constraint FKSEGUITO_FK
     foreign key (Username)
     references UTENTE;

alter table SEGUIRE add constraint FKSEGUE
     foreign key (SEG_Username)
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
     foreign key (RES_IdMembro)
     references AMMINISTRATORE;

--Not implemented
--alter table UTENTE add constraint ID_UTENTE_CHK
--     check(exists(select * from STORICO_PASSWORD
--                  where STORICO_PASSWORD.Username = Username)); 

alter table UTENTE add constraint FKDOMICILIO_FK
     foreign key (IdRegione)
     references REGIONE;


-- Index Section
-- _____________ 

create index FKLUOGO_PUBBLICAZIONE_IND
     on CONTENUTO (IdRegione);

create index FKRISPOSTA_IND
     on CONTENUTO (A_Username, A_IdContenuto);

create index FKIMPERSONAZIONE_IND
     on MEMBRO (Username);

create index FKAGGIUNTO_IND
     on MEMBRO (Amministratore);

create index FKMITTENTE_IND
     on MESSAGGIO (IdMembro);

create index FKCITAZIONE_IND
     on MESSAGGIO (Citazione);

create index FKRIFERIMENTO_IND
     on MESSAGGIO (Username, IdContenuto);

create index FKELIMINAZIONE_IND
     on MESSAGGIO (AmministratoreEliminatore);

create index FKREA_CON_IND
     on REAZIONE (R_C_Username, R_C_IdContenuto);

create index FKPARTE_DI_IND
     on REGIONE (Superregione);

create index FKSEGUITO_IND
     on SEGUIRE (Username);

create index FKRESPONSABILE_IND
     on USCITA (RES_IdMembro);

create index FKDOMICILIO_IND
     on UTENTE (IdRegione);

