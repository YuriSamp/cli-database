package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type Entry struct {
	value string
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

func (db *Database) Copy(source string, destination string) string {
	layer := db.getcurrLayer()
	v, ok := layer[source]

	if !ok {
		return "ERR source not found"
	}

	db.Set(destination, v.value)

	return "OK"
}

func (db *Database) Delete(key string) string {
	layer := db.getcurrLayer()
	_, ok := layer[key]

	if !ok {
		return "ERR Key not found"
	}

	delete(layer, key)

	return "OK"
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

func (db *Database) List() []string {
	entries := []string{}

	layer := db.getcurrLayer()

	for key, entry := range layer {
		entries = append(entries, fmt.Sprintf("key: %s, value: %s", key, entry.value))
	}

	return entries
}

func (db *Database) TTL(key string) string {
	layer := db.getcurrLayer()
	v, ok := layer[key]

	if !ok {
		return "-2"
	}

	if v.ttl == -1 {
		return "-1"
	}

	return fmt.Sprintf("%d", v.ttl)
}

func (db *Database) Persist(key string) string {
	layer := db.getcurrLayer()
	v, ok := layer[key]

	if !ok || v.ttl == -1 {
		return "0"
	}

	newEntry := &Entry{value: v.value, ttl: -1}
	layer[key] = *newEntry

	return "1"
}

// func (db *Database) Expire(key string, time string) string {
// 	parsedTime, err := strconv.Atoi(time)

// 	if err != nil {
// 		return err.Error()
// 	}

// 	layer := db.getcurrLayer()
// 	v, ok := layer[key]

// 	if !ok {
// 		return "0"
// 	}

// 	if parsedTime < 0 {
// 		db.Delete(key)
// 		return "1"
// 	}

// 	newEntry := &Entry{value: v.value, ttl: parsedTime}
// 	layer[key] = *newEntry

// 	return "1"
// }

func (db *Database) getcurrLayer() map[string]Entry {
	return db.dbLayers[db.pointer]
}

func (db *Database) GetFirstLayer() map[string]Entry {
	return db.dbLayers[0]
}

func (db *Database) popLastLayer() {
	db.dbLayers = db.dbLayers[0 : len(db.dbLayers)-1]
}
