package session

import "testing"

type User struct {
	Name string `goorm:"primary key"`
	Age int
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	s.DropTable()
	s.CreateTable()
	if !s.HasTable(){
		t.Fatal("failed to create user")
	}
}

func TestSession_Model(t *testing.T) {
	s := NewSession().Model(&User{})
	table := s.RefTable()
	s.Model(&Session{})
	if table.Name != "User" || s.RefTable().Name!="Session"{
		t.Fatal("failed to change model")
	}
}
