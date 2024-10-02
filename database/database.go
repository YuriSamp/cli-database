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

func (db *Database) BeginTransaction() {
	newLayer := make(map[string]string)
	db.dbLayers = append(db.dbLayers, newLayer)
	db.pointer += 1
	fmt.Println(db.pointer)
}

func (db *Database) Rollback() error {
	if db.pointer == 0 {
		return fmt.Errorf("ERR Invalid command when outside a transaction")
	}

	db.popLastLayer()
	db.pointer -= 1
	fmt.Println(db.pointer)

	return nil
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

func (db *Database) Delete(key string) {
	layer := db.getcurrLayer()
	_, ok := layer[key]

	if !ok {
		fmt.Println("ERR Key not found")
		return
	}

	delete(layer, key)

	fmt.Println("OK")
}

func (db *Database) Commit() error {
	if db.pointer == 0 {
		return fmt.Errorf("ERR Invalid command when outside a transaction")
	}

	topLayer := db.getcurrLayer()
	db.pointer -= 1

	currLayer := db.getcurrLayer()

	for k, v := range topLayer {
		currLayer[k] = v
	}

	db.popLastLayer()

	return nil
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
