package database

import "fmt"

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

	topLayer := db.getCurrLayer()
	db.pointer -= 1

	currLayer := db.getCurrLayer()

	for k, v := range topLayer {
		currLayer[k] = v
	}

	db.popLastLayer()

	return "OK"
}
