package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/icrowley/fake"
	_ "github.com/lib/pq"
	"github.com/michelececcacci/db-proj/model"
	"github.com/michelececcacci/db-proj/queries"

	"time"
)

const (
	dbname   = "postgres"
	user     = "postgres"
	password = "postgres"
	host     = "localhost"
	port     = "55432"
)

func run() error {
	s := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", dbname, user, password, host, port)
	db, err := sql.Open("postgres", s)
	if err != nil {
		fmt.Println("main", err)
		return err
	}
	model := model.New(db, context.Background())
	populateUsers(model, 100)
	populateFollowers(model, 500)
	populateChats(model, 50)
	populateMembers(model, 20, true)
	populateMembers(model, 7000, false)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func fakeDate() time.Time {
	return time.Date(fake.Year(2021, 2022), time.Month(fake.MonthNum()), fake.Day(), 0, 0, 0, 0, time.UTC)
}

func populateUsers(m model.Model, n int) {
	for i := 0; i < n; {
		u := queries.InsertUserParams{
			Username:      fake.UserName(),
			Nome:          sql.NullString{String: fake.FirstName(), Valid: true},
			Cognome:       sql.NullString{String: fake.LastName(), Valid: true},
			Datadinascita: sql.NullTime{Time: fakeDate(), Valid: true},
		}
		fmt.Println(u)
		m.InsertUser(u)
		i++
	}
}

func populateFollowers(m model.Model, n int) {
	followings, _ := m.GetAllPossibleFollowings()
	k := 0
	i := 0
	for ; i < n && k < len(followings); k++ {
		dataInizio := fakeDate()
		seguace, seguito := followings[k].U1.String, followings[k].U2.String
		f := queries.InsertFollowerParams{
			Usernameseguace: seguace,
			Usernameseguito: seguito,
			Datainizio:      dataInizio,
			Datafine:        sql.NullTime{Time: dataInizio.AddDate(0, 1, 0), Valid: true},
		}
		fmt.Println(f)
		err := m.InsertFollower(f)
		if err == nil {
			i++
		} else {
			fmt.Println(err)
		}
	}
	if i < n {
		fmt.Println("We couldn't create all followings")
	}
}

func populateChats(m model.Model, n int) {
	for i := 0; i < n; i++ {
		c := queries.InsertChatParams{
			Nome:        fake.Word(),
			Descrizione: fake.WordsN(5),
		}
		creator, _ := m.GetRandomUser()
		fmt.Println(c)
		m.InsertChat(c, creator)
	}
}

func populateMembers(m model.Model, n int, asAdmin bool) {
	members, _ := m.GetAllPossibleMembers()
	i := 0
	k := 0
	for ; i < n && k < len(members); k++ {
		u, chat := members[k].Username.String, members[k].Idchat.Int32
		admin, err := m.GetRandomAdminInChat(chat)
		if err != nil {
			fmt.Println("Members3", err, admin)
		}
		if asAdmin {
			_, err = m.InsertAdminMember(u, chat, admin)
		} else {
			_, err = m.InsertNotAdminMember(u, chat, admin)
		}
		if err == nil {
			i++
		} else {
			// fmt.Println("Members4", err)
		}
	}
	if k == len(members) && i < n {
		fmt.Println("We couldn't create all members")
	}
}
