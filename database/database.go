package database

type Database struct {
	dbLayers []map[string]string
	pointer  int
}

func New() *Database {
	layers := []map[string]string{}
	db := &Database{dbLayers: layers, pointer: 0}
	return db
}
