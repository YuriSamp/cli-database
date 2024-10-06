package database

import "testing"

func TestDeleteCommand(t *testing.T) {
	db := New()
	msg := db.Delete("teste")
	expected := "ERR Key not found"

	if msg != expected {
		t.Errorf("Database delete command don't return the right value. Expected %s, got=%s", expected, msg)
	}

	db.Set("teste", "1")

	msg = db.Delete("teste")

	if msg != "OK" {
		t.Errorf("Database delete command don't return the right value. Expected %s, got=%s", "OK", msg)
	}

	msg = db.Get("teste")

	if msg != "NIL" {
		t.Errorf("Database don't deleted the key")
	}
}

func TestBeginCommand(t *testing.T) {
	db := New()

	if db.pointer != 0 {
		t.Errorf("Database initialized with wrong pointer")
	}

	db.Set("teste", "1")
	db.Set("batata", "2")
	db.Set("carro", "3")

	newPointer := db.BeginTransaction()
	currLayer := db.getcurrLayer()

	if len(currLayer) != 0 {
		t.Errorf("Database failed to initialize a new empty layer")
	}

	if newPointer != "1" {
		t.Errorf("Database set the wrong pointer")
	}
}
