package database

import (
	"strconv"
)

func (db *Database) Set(key, value string) string {
	layer := db.getcurrLayer()

	newEntry := &Entry{value: value, ttl: -1}

	layer[key] = *newEntry

	return "OK"
}

func (db *Database) Get(key string) string {
	layer := db.getcurrLayer()
	v, ok := layer[key]

	if ok {
		return v.value
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

		valuesToPrint = append(valuesToPrint, v.value)
	}

	return valuesToPrint
}

func (db *Database) Mset(keyValues []string) string {

	for i := 0; i < len(keyValues); i += 2 {
		db.Set(keyValues[i], keyValues[i+1])
	}

	return "OK"
}

func (db *Database) Incr(key string) string {
	v := db.Get(key)

	if v == "NIL" {
		db.Set(key, "0")
		return "OK"
	}

	for _, ch := range v {
		if !isDigit(ch) {
			return "ERR this key is not a number"
		}
	}

	value, _ := strconv.Atoi(v)
	parsedInt := strconv.Itoa(value + 1)
	db.Set(key, parsedInt)

	return "OK"
}

func (db *Database) Decr(key string) string {

	v := db.Get(key)

	for _, ch := range v {
		if !isDigit(ch) {
			return "ERR this key is not a number"
		}
	}

	value, _ := strconv.Atoi(v)
	parsedInt := strconv.Itoa(value - 1)
	db.Set(key, parsedInt)

	return "OK"
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}
