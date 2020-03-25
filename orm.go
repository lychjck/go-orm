package go_orm

import (
	"go-orm/log"
	"go-orm/session"
	"database/sql"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver,source string) (e *Engine,err error)  {
	db,err := sql.Open(driver,source)
	if err!=nil{
		log.Error(err)
		return
	}
	if err = db.Ping();err != nil{
		log.Error(err)
		return
	}
	e = &Engine{db:db}
	log.Info("Connect Successfully!")
	return
}

func (engine *Engine)Close()  {
	if err := engine.db.Close();err != nil{
		log.Error("Failed to close the database!")
	}
	log.Info("Close Successfully!")
}

func (engine *Engine)NewSession()  *session.Session {
	return session.New(engine.db)
}