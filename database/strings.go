package database

import (
	"strconv"
)

func (db *Database) Set(key string, value interface{}) string {
	layer := db.getCurrLayer()

	newEntry := &Entry{value: value, ttl: -1}

	layer[key] = *newEntry

	return "OK"
}

func (db *Database) Get(key string) string {
	layer := db.getCurrLayer()
	v, ok := layer[key]

	if !ok {
		return "NIL"
	}

	switch value := v.value.(type) {
	case string:
		return value
	default:
		return "NIL"
	}

}

func (db *Database) Mget(keys []string) []string {

	layer := db.getCurrLayer()

	valuesToPrint := []string{}

	for _, key := range keys {

		v, ok := layer[key]

		if !ok {
			valuesToPrint = append(valuesToPrint, "NIL")
			continue
		}

		if value, ok := v.value.(string); ok {
			valuesToPrint = append(valuesToPrint, value)
		} else {
			valuesToPrint = append(valuesToPrint, "NIL")
		}

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
