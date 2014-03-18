package rethink

import (
	"fmt"
	r "github.com/dancannon/gorethink"
)

// DB holds the connection to the database.
type DB struct {
	Session *r.Session
}

// NewDB creates a new database.
func NewDB(session *r.Session) *DB {
	return &DB{
		Session: session,
	}
}

var emptyMap map[string]interface{}

// Get retrieves a single item from the database and returns it as a map.
func (db *DB) Get(table, id string) (map[string]interface{}, error) {
	result, err := r.Table(table).Get(id).RunRow(db.Session)
	switch {
	case err != nil:
		return emptyMap, err
	case result.IsNil():
		return emptyMap, fmt.Errorf("no such id: %s", id)
	default:
		var response map[string]interface{}
		result.Scan(&response)
		return response, nil
	}
}

// GetAll retrieves multiple items from the database and returns them as a
// list of maps.
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
