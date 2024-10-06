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
