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

func (db *Database) Set(key string, value string) {
	msg := db.hasKey(key, value)

	layer := db.getLayer()
	layer[key] = value

	fmt.Println(msg)
}

func (db *Database) Get(key string) {
	layer := db.getLayer()
	v, ok := layer[key]

	if ok {
		fmt.Println(v)
	}

	fmt.Println("NIL")
}

func (db *Database) Commit() error {
	if db.pointer == 0 {
		return fmt.Errorf("ERR Invalid command when outside a transaction")
	}

	topLayer := db.getLayer()
	db.pointer -= 1

	currLayer := db.getLayer()

	for k, v := range topLayer {
		currLayer[k] = v
	}

	db.popLastLayer()

	return nil
}

func (db *Database) hasKey(key string, value string) string {
	layer := db.getLayer()
	_, ok := layer[key]

	if ok {
		return fmt.Sprintf("TRUE %s", value)
	}

	return fmt.Sprintf("FALSE %s", value)
}

func (db *Database) getLayer() map[string]string {
	return db.dbLayers[db.pointer]
}

func (db *Database) popLastLayer() {
	db.dbLayers = db.dbLayers[0 : len(db.dbLayers)-1]
}
