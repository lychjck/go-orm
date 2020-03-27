package dialect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSqlite3_DataTypeOf(t *testing.T) {
	dial := &sqlite3{}
	fmt.Println(reflect.TypeOf(dial))
	cases := []struct{
		Value interface{}
		Type string
	}{
		{"tom","text"},
		{123,"integer"},
		{1.2,"real"},
		{[]int{1,2,3},"blob"},
	}

	for _,c := range cases{
		fmt.Println(reflect.TypeOf(dial))
		if typ := dial.DataTypeOf(reflect.ValueOf(c.Value)); typ != c.Type{
			t.Fatalf("except %s,but got %s",c.Type,typ)
		}
	}
}
