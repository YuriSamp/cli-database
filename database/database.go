package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type Database struct {
	dbLayers []map[string]string
	pointer  int
}

func New() *Database {
	layers := make([]map[string]string, 1)

	var initialLayer map[string]string

	data, err := os.ReadFile("./database.json")

	if err != nil {
		initialLayer = make(map[string]string)
	} else {
		json.Unmarshal(data, &initialLayer)
	}

	layers[0] = initialLayer
	db := &Database{dbLayers: layers, pointer: 0}
	return db
}

func (db *Database) BeginTransaction() string {
	newLayer := make(map[string]string)
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

func (db *Database) Set(key string, value string) string {
	msg := db.hasKey(key, value)

	layer := db.getcurrLayer()
	layer[key] = value

	return msg
}

func (db *Database) Get(key string) string {
	layer := db.getcurrLayer()
	v, ok := layer[key]

	if ok {
		return v
	}

	return "NIL"
}

func (db *Database) Mget(keys []string) []string {

	layer := db.getcurrLayer()

	valuesToPrint := []string{}

	for _, key := range keys {

		v, ok := layer[key]

		if !ok {
			valuesToPrint = append(valuesToPrint, "NIL")
			continue
		}

		valuesToPrint = append(valuesToPrint, v)
	}

	return valuesToPrint
}

func (db *Database) Mset(keyValues []string) string {

	for i := 0; i < len(keyValues); i += 2 {
		db.Set(keyValues[i], keyValues[i+1])
	}

	return "OK"
}

func (db *Database) Copy(source string, destination string) string {
	layer := db.getcurrLayer()
	v, ok := layer[source]

	if !ok {
		return "ERR source not found"
	}

	db.Set(destination, v)

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

	for key, value := range layer {
		entries = append(entries, fmt.Sprintf("key: %s, value: %s", key, value))
	}

	return entries
}

func (db *Database) hasKey(key string, value string) string {
	layer := db.getcurrLayer()
	_, ok := layer[key]

	if ok {
		return fmt.Sprintf("TRUE %s", value)
	}

	return fmt.Sprintf("FALSE %s", value)
}

func (db *Database) getcurrLayer() map[string]string {
	return db.dbLayers[db.pointer]
}

func (db *Database) GetFirstLayer() map[string]string {
	return db.dbLayers[0]
}

func (db *Database) popLastLayer() {
	db.dbLayers = db.dbLayers[0 : len(db.dbLayers)-1]
}
