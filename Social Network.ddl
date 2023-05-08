-- *********************************************
-- * SQL PostgreSQL generation                 
-- *--------------------------------------------
-- * DB-MAIN version: 11.0.2              
-- * Generator date: Sep 20 2021              
-- * Generation date: Thu May  4 16:02:22 2023 
-- * LUN file: /home/luca/db-proj/Social Network.lun 
-- * Schema: Logica/1 
-- ********************************************* 


-- Database Section
-- ________________ 

create database Logica;


-- Tables Section
-- _____________ 

create table CHAT (
     Nome varchar(100) not null,
     IdChat numeric(1) not null,
     constraint IDCHAT primary key (IdChat));

create table CONTENUTO (
     Username varchar(20) not null,
     Testo varchar(1000) not null,
     IdContenuto numeric(1) not null,
     TimestampPubblicazione date not null,
     constraint IDCONTENUTO primary key (Username, IdContenuto));

create table MEMBRO (
     TimestampUltimaLettura date not null,
     Amministratore char not null,
     constraint IDMEMBRO primary key (, ));

create table MESSAGGIO (
     Username varchar(20) not null,
     IdMessaggio numeric(1) not null,
     Testo varchar(500) not null,
     TimestampInvio date not null,
     IdChat numeric(1) not null,
     constraint IDMESSAGGIO primary key (Username, IdMessaggio));

create table REAZIONE (
     LikeDislike numeric(1) not null,
     Timestamp char(1) not null,
     constraint IDREAZIONE primary key (, ));

create table REGIONE (
     IdRegione numeric(1) not null,
     Nome varchar(50) not null,
     CONTIENE_ numeric(1),
     constraint IDREGIONE primary key (IdRegione));

create table SEGUIRE (
,
     constraint IDSEGUIRE primary key (, ));

create table STORICO_PASSWORD (
     DataInserimento date not null,
     Password varchar(20) not null,
     Username varchar(20) not null,
     constraint IDSTORICO_PASSWORD primary key (Username, Password),
     constraint IDSTORICO_PASSWORD_1 unique (Username, DataInserimento));

create table UTENTE (
     Username varchar(20) not null,
     Generalità -- Compound attribute -- not null,
     IdRegione numeric(1),
     constraint IDUTENTE_ID primary key (Username));


-- Constraints Section
-- ___________________ 

alter table CONTENUTO add constraint FKPUBBLICAZIONE
     foreign key (Username)
     references UTENTE;

alter table MESSAGGIO add constraint FKSEQUENZA
     foreign key (IdChat)
     references CHAT;

alter table MESSAGGIO add constraint FKMITTENTE
     foreign key (Username)
     references UTENTE;

alter table REGIONE add constraint FKPARTE_DI
     foreign key (CONTIENE_)
     references REGIONE;

alter table STORICO_PASSWORD add constraint FKSCELTA
     foreign key (Username)
     references UTENTE;

--Not implemented
--alter table UTENTE add constraint IDUTENTE_CHK
--     check(exists(select * from STORICO_PASSWORD
--                  where STORICO_PASSWORD.Username = Username)); 

alter table UTENTE add constraint FKDOMICILIO
     foreign key (IdRegione)
     references REGIONE;


-- Index Section
-- _____________ 

