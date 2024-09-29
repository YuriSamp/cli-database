package database

import "fmt"

type Database struct {
	dbLayers []map[string]string
	pointer  int
}

func New() *Database {
	layers := make([]map[string]string, 1)
	layers[0] = make(map[string]string)
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

func (db *Database) Set(key string, value string) string {
	msg := db.hasKey(key, value)

	layer := db.dbLayers[db.pointer]
	layer[key] = value
	return msg
}

func (db *Database) Get(key string) string {
	layer := db.dbLayers[db.pointer]
	v, ok := layer[key]

	if ok {
		return v
	}

	return "NIL"
}

func (db *Database) hasKey(key string, value string) string {
	layer := db.dbLayers[db.pointer]
	_, ok := layer[key]

	if ok {
		return fmt.Sprintf("TRUE %s", value)
	}

	return fmt.Sprintf("FALSE %s", value)
}
