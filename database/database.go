package database

import (
	"encoding/json"
	"os"
)

type Entry struct {
	value interface{}
	ttl   int
}

type Database struct {
	dbLayers []map[string]Entry
	pointer  int
}

func New() *Database {
	layers := make([]map[string]Entry, 1)

	var initialLayer map[string]Entry

	data, err := os.ReadFile("./database.json")

	if err != nil {
		initialLayer = make(map[string]Entry)
	} else {
		json.Unmarshal(data, &initialLayer)
	}

	layers[0] = initialLayer
	db := &Database{dbLayers: layers, pointer: 0}
	return db
}

func (db *Database) getCurrLayer() map[string]Entry {
	return db.dbLayers[db.pointer]
}

func (db *Database) GetFirstLayer() map[string]Entry {
	return db.dbLayers[0]
}

func (db *Database) popLastLayer() {
	db.dbLayers = db.dbLayers[0 : len(db.dbLayers)-1]
}
