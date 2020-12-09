package model

import "database/sql"

var db DB

//DB is a inteface for db
type DB interface {
	QueryRow(string, ...interface{}) Row
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (Result, error)
}

//Row is a intrefave for row
type Row interface {
	Scan(...interface{}) error
}

//Result is a intrefave for result
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type sqlDB struct {
	db *sql.DB
}

func (s sqlDB) QueryRow(query string, args ...interface{}) Row {
	return s.db.QueryRow(query, args...)
}

func (s sqlDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}

func (s sqlDB) Exec(query string, args ...interface{}) (Result, error) {
	return s.db.Exec(query, args...)
}

//SetDatabase function set db
func SetDatabase(database *sql.DB) {
	db = &sqlDB{database}
}
