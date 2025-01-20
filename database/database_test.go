package database

import (
	"testing"
)

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

func TestRollbackCommand(t *testing.T) {
	db := New()

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

	newPointer = db.Rollback()
	currLayer = db.getcurrLayer()

	if newPointer != "0" {
		t.Errorf("Database set the wrong pointer at rollback command")
	}

	if len(currLayer) != 3 {
		t.Errorf("Database rollback to a empty layer")
	}
}

func TestRollbackCommandOutsideATransaction(t *testing.T) {
	db := New()

	err := db.Rollback()

	if err != "ERR Invalid command when outside a transaction" {
		t.Errorf("Something got wrong, rollback ran outside a transaction")
	}
}

func TestCopyCommand(t *testing.T) {
	db := New()

	setValue := "10"
	db.Set("teste", "10")
	msg := db.Copy("carro", "aviao")
	expectedError := "ERR source not found"

	if msg != expectedError {
		t.Errorf("Database should throw an error if an unknown source is gave")
	}

	db.Copy("teste", "teste2")
	value := db.Get("teste2")

	if value != setValue {
		t.Errorf("Database failed to copy a value. Expected %s, got=%s", setValue, value)
	}
}

func TestCommitCommand(t *testing.T) {
	db := New()

	db.Set("teste", "1")
	db.Set("batata", "2")
	db.Set("carro", "3")

	db.BeginTransaction()

	db.Set("aviao", "4")
	db.Set("onibus", "5")
	db.Set("trem", "6")

	db.Commit()

	layer := db.getcurrLayer()

	if len(layer) != 6 {
		t.Errorf("Wrong length of database after commit. Expeted %d, got=%d", 6, len(layer))
	}

	value := db.Get("aviao")

	if value != "4" {
		t.Errorf("Wrong value after commit. Expected %s, got=%s", "4", value)
	}
}

func TestCommitCommandOutsideATransaction(t *testing.T) {
	db := New()

	err := db.Commit()

	if err != "ERR Invalid command when outside a transaction" {
		t.Errorf("Something got wrong, commit ran outside a transaction")
	}
}
