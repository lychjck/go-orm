package go_orm

import (
	"database/sql"
	"go-orm/dialect"
	"go-orm/log"
	"go-orm/session"
)

type Engine struct {
	db *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	dial,ok := dialect.GetDialect(driver)
	if !ok{
		log.Errorf("dialect %s not found",driver)
		return
	}
	e = &Engine{db: db,dialect:dial}
	log.Info("Connect Successfully!")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close the database!")
	}
	log.Info("Close Successfully!")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db,engine.dialect)
}
