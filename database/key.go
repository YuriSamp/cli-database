package database

import (
	"fmt"
	"strconv"
)

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

func (db *Database) Expire(key, time string) string {
	parsedTime, err := strconv.Atoi(time)

	if err != nil {
		return err.Error()
	}

	layer := db.getcurrLayer()
	v, ok := layer[key]

	if !ok {
		return "0"
	}

	if parsedTime < 0 {
		db.Delete(key)
		return "1"
	}

	newEntry := &Entry{value: v.value, ttl: parsedTime}
	layer[key] = *newEntry

	return "1"
}

func (db *Database) Copy(source, destination string) string {
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

func (db *Database) Rename(sorce, desination string) string {
	layer := db.getcurrLayer()
	v, ok := layer[sorce]

	if !ok {
		return "ERR Key not found"
	}

	layer[desination] = v
	db.Delete(sorce)

	return "OK"
}

func (db *Database) Exists(keys []string) string {

	keyCount := 0
	layer := db.getcurrLayer()

	for _, key := range keys {
		_, ok := layer[key]

		if ok {
			keyCount++
		}
	}

	return fmt.Sprintf("%d", keyCount)
}
