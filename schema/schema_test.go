package schema

import (
	"fmt"
	"go-orm/dialect"
	"reflect"
	"testing"
)

type User struct {
	Name string `goorm:"primary key"`
	Age int
}

var TestDial,_ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{},TestDial)
	fmt.Println(reflect.TypeOf(schema))
	if schema.Name != "User" || len(schema.Fields) != 2{
		t.Fatal("failed to parse user struct")
	}
	if schema.GetField("Name").Tag != "primary key"{
		t.Fatal("failed to parse primary key")
	}
}
