package repository

import "testing"

func TestInitDB(t *testing.T) {
	InitDB()
	if closed {
		t.Errorf("Error: %v", "Database is closed")
	}
}

func TestDisposeDB(t *testing.T) {
	DisposeDB()
	if !closed {
		t.Errorf("Error: %v", "Database is not closed")
	}
}
