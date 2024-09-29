package database

import "fmt"

type Database struct {
	dbLayers []map[string]string
	pointer  int
}

func New() *Database {
	layers := make([]map[string]string, 1)
	db := &Database{dbLayers: layers, pointer: 0}
	return db
}

func (db *Database) BeginTransaction() {
	newLayer := make(map[string]string)
	db.dbLayers = append(db.dbLayers, newLayer)
	db.pointer += 1
}

func (db *Database) Rollback() error {
	if db.pointer == 0 {
		return fmt.Errorf("ERR Invalid command when outside a transaction")
	}

	db.dbLayers = db.dbLayers[0 : len(db.dbLayers)-1]
	db.pointer -= 1
	return nil
}
