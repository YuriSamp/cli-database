package database

import (
	"encoding/json"
	"fmt"
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

func (db *Database) BeginTransaction() string {
	newLayer := make(map[string]Entry)
	db.dbLayers = append(db.dbLayers, newLayer)
	db.pointer += 1

	return fmt.Sprintf("%d", db.pointer)
}

func (db *Database) Rollback() string {
	if db.pointer == 0 {
		return "ERR Invalid command when outside a transaction"
	}

	db.popLastLayer()
	db.pointer -= 1

	return fmt.Sprintf("%d", db.pointer)
}

func (db *Database) Commit() string {
	if db.pointer == 0 {
		return "ERR Invalid command when outside a transaction"
	}

	topLayer := db.getcurrLayer()
	db.pointer -= 1

	currLayer := db.getcurrLayer()

	for k, v := range topLayer {
		currLayer[k] = v
	}

	db.popLastLayer()

	return "OK"
}

func (db *Database) getcurrLayer() map[string]Entry {
	return db.dbLayers[db.pointer]
}

func (db *Database) GetFirstLayer() map[string]Entry {
	return db.dbLayers[0]
}

func (db *Database) popLastLayer() {
	db.dbLayers = db.dbLayers[0 : len(db.dbLayers)-1]
}
