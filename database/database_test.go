package database

import "testing"

func TestDatabaseGetCommandWithValue(t *testing.T) {

	db := New()
	expected := "1"

	db.Set("teste", expected)
	msg := db.Get("teste")

	if msg != expected {
		t.Errorf("Database get command don't return the right value. Expected %s, got=%s", msg, expected)
	}
}

func TestDatabaseGetCommandWithoutValue(t *testing.T) {

	db := New()
	expected := "NIL"
	msg := db.Get("teste")

	if msg != expected {
		t.Errorf("Database get command don't return the right value. Expected %s, got=%s", msg, expected)
	}
}
