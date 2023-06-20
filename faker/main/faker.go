package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/icrowley/fake"

	"github.com/brianvoe/gofakeit/v6"
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

const (
	chats    = 50
	messages = 2
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
	populateChats(model, chats)
	populateMembers(model, 20, true)
	populateMembers(model, 7000, false)
	exitMembers(model, 100, true)
	exitMembers(model, 50, false)
	populateMessages(model, chats, messages)
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

func populateMembers(m model.Model, n int, asAdmin bool) error {
	members, _ := m.GetAllPossibleMembers()
	i := 0
	k := 0
	for ; i < n && k < len(members); k++ {
		u, chat := members[k].Username.String, members[k].Idchat.Int32
		admin, err := m.GetRandomAdminInChat(chat)
		if err != nil {
			return fmt.Errorf("members3\n%s\n%v", err, admin)
		}
		if asAdmin {
			_, err = m.InsertAdminMember(u, chat, admin)
		} else {
			_, err = m.InsertNotAdminMember(u, chat, admin)
		}
		if err == nil {
			i++
		}
	}
	if k == len(members) && i < n {
		return fmt.Errorf("couldn't create all members")
	}
	return nil
}

func exitMembers(m model.Model, n int, exile bool) error {
	for i := 0; i < n; i++ {
		chat, err := m.GetRandomChat()
		if err != nil {
			return err
		}
		sqlMember, sqlDate, err := m.GetRandomCurrentMemberInChat(chat)
		if err != nil {
			return err
		}
		member := sqlMember.Int32
		joinDate := sqlDate.Time
		admin, err := m.GetRandomAdminInChat(chat)
		if err != nil {
			return err
		}
		err = m.ExitMember(member,
			gofakeit.DateRange(joinDate, time.Now()),
			sql.NullInt32{
				Int32: admin,
				Valid: exile,
			},
			sql.NullString{
				String: gofakeit.Quote(),
				Valid:  true,
			})
		if err != nil {
			return err
		}
	}
	return nil
}

func populateMessages(m model.Model, chats, messages int) error {
	for i := 0; i < chats; i++ {
		chat, err := m.GetRandomChat()
		if err != nil {
			return err
		}
		members, _ := m.GetChatMembers(chat)
		var membersInt []int
		for _, member := range members {
			membersInt = append(membersInt, int(member))
		}
		for j := 0; j < messages; j++ {
			m.InsertMessage(queries.InsertMessageParams{
				Testo:          gofakeit.Quote(),
				Timestampinvio: fakeDate(),
				Mittente:       int32(gofakeit.RandomInt(membersInt)),
			})
		}
	}
	return nil
}
