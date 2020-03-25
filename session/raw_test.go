package session

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

var TestDB *sql.DB

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "../test.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(TestDB)
}

func TestSession_Exec(t *testing.T) {
	s := NewSession()
	s.Raw("drop table if exists user;").Exec()
	s.Raw("create table user(name text);").Exec()
	result, _ := s.Raw("insert into user(`name`) values (?),(?)", "Dog", "cat").Exec()
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatal("插入数据不对")
	}
}

func TestSession_QueryRows(t *testing.T) {
	s := NewSession()
	s.Raw("drop table  if exists user;").Exec()
	s.Raw("create table user(name text);").Exec()
	row := s.Raw("select count(*) from user").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil || count != 0 {
		fmt.Println(count)
		fmt.Println(err)
		t.Fatal(count)
		t.Fatal("failed to query db")
	}
}
