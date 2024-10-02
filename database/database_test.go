package database

import "testing"

func TestGetCommandWithValue(t *testing.T) {

	db := New()
	expected := "1"

	db.Set("teste", expected)
	msg := db.Get("teste")

	if msg != expected {
		t.Errorf("Database get command don't return the right value. Expected %s, got=%s", msg, expected)
	}
}

func TestGetCommandWithoutValue(t *testing.T) {

	db := New()
	expected := "NIL"
	msg := db.Get("teste")

	if msg != expected {
		t.Errorf("Database get command don't return the right value. Expected %s, got=%s", msg, expected)
	}
}

func TestSetCommand(t *testing.T) {

	db := New()
	msg := db.Set("teste", "1")
	expected := "FALSE 1"

	if msg != expected {
		t.Errorf("Database set command don't return the right value. Expected %s, got=%s", msg, expected)
	}

	msg = db.Set("teste", "1")
	expected = "TRUE 1"

	if msg != expected {
		t.Errorf("Database set command don't return the right value. Expected %s, got=%s", msg, expected)
	}
}

