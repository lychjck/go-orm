package schema

import (
	"go-orm/dialect"
	"testing"
)

type User struct {
	Name string `goorm:"primary key"`
	Age int
}

var TestDial,_ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{},TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2{
		t.Fatal("failed to parse user struct")
	}
	if schema.GetField("Name").Tag != "primary key"{
		t.Fatal("failed to parse primary key")
	}
}
