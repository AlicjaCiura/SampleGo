package model

import "database/sql"

type mockDB struct {
	lastQuery   string
	lastArgs    []interface{}
	returnedRow Row
}

func (db *mockDB) QueryRow(query string, args ...interface{}) Row {
	db.lastQuery = query
	db.lastArgs = args
	return db.returnedRow
}

func (db *mockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (db *mockDB) Exec(query string, args ...interface{}) (Result, error) {
	return nil, nil
}

type mockRow struct{}

func (m mockRow) Scan(...interface{}) error {
	return nil
}
