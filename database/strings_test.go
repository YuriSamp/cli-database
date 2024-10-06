package database

import "testing"

func TestGetCommandWithValue(t *testing.T) {

	db := New()
	expected := "1"

	db.Set("teste", expected)
	value := db.Get("teste")

	if value != expected {
		t.Errorf("Database get command don't return the right value. Expected %s, got=%s", expected, value)
	}
}

func TestGetCommandWithoutValue(t *testing.T) {

	db := New()
	expected := "NIL"
	value := db.Get("teste")

	if value != expected {
		t.Errorf("Database get command don't return the right value. Expected %s, got=%s", expected, value)
	}
}

func TestSetCommand(t *testing.T) {

	db := New()
	valueToBeSet := "1"
	msg := db.Set("teste", valueToBeSet)
	expected := "OK"

	if msg != "OK" {
		t.Errorf("Database set command don't return the right value. Expected %s, got=%s", expected, msg)
	}

	value := db.Get("teste")

	if value != valueToBeSet {
		t.Errorf("Database don't seted the right value. Expected=%s, got=%s", expected, value)
	}
}
