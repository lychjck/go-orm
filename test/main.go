package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go-orm"
)

func main() {
	engine, _ := go_orm.NewEngine("sqlite3", "test.db")
	defer engine.Close()
	s := engine.NewSession()
	s.Raw("drop table if exist user;").Exec()
	s.Raw("create table user(name text);").Exec()
	s.Raw("create table user(name text);").Exec()
	result, _ := s.Raw("insert into user(`name`) values (?),(?)", "Dog", "Cat").Exec()
	count, _ := result.RowsAffected()
	fmt.Println(count)
}
