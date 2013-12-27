package rethink

import (
	r "github.com/dancannon/gorethink"
)

type DB struct {
	Session *r.Session
}

func NewDB(session *r.Session) *DB {
	return &DB{
		Session: session,
	}
}

func (db *DB) Get(table, id string) (map[string]interface{}, error) {
	result, err := r.Table(table).Get(id).RunRow(db.Session)
	if err != nil || result.IsNil() {
		return nil, err
	}

	var response map[string]interface{}
	result.Scan(&response)
	return response, nil
}

func (db *DB) GetAll(query r.RqlTerm) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	rows, err := query.Run(db.Session)
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var response map[string]interface{}
		rows.Scan(&response)
		results = append(results, response)
	}

	return results, nil
}
