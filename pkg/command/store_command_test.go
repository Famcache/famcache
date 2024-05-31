package command

import (
	"famcache/domain/command"
	"testing"
)

func TestNewCommand_Set(t *testing.T) {
	query := "1 SET key value 3000"
	q := NewStoreCommand(command.CommandSet, query)

	if q == nil {
		t.Error("Query is nil")
	}

	if q.ID() != "1" {
		t.Errorf("Expected ID to be 1, got %s", q.ID())
	}

	if q.Type() != command.CommandSet {
		t.Errorf("Expected type to be SET, got %s", q.Type())
	}

	if q.Key() != "key" {
		t.Errorf("Expected key to be key, got %s", q.Key())
	}

	if q.Value() == nil {
		t.Error("Value is nil")
	}

	if *q.Value() != "value" {
		t.Errorf("Expected value to be value, got %s", *q.Value())
	}

	if q.TTL() == nil {
		t.Error("TTL is nil")
	}

	if *q.TTL() != 3000 {
		t.Errorf("Expected TTL to be 5, got %d", *q.TTL())
	}
}

func TestNewCommand_Get(t *testing.T) {
	query := "1 GET key"
	q := NewStoreCommand(command.CommandGet, query)

	if q == nil {
		t.Error("Query is nil")
	}

	if q.ID() != "1" {
		t.Errorf("Expected ID to be 1, got %s", q.ID())
	}

	if q.Type() != command.CommandGet {
		t.Errorf("Expected type to be GET, got %s", q.Type())
	}

	if q.Key() != "key" {
		t.Errorf("Expected key to be key, got %s", q.Key())
	}

	if q.Value() != nil {
		t.Error("Value is not nil")
	}

	if q.TTL() != nil {
		t.Error("TTL is not nil")
	}
}

func TestNewCommand_Delete(t *testing.T) {
	query := "1 DELETE key"
	q := NewStoreCommand(command.CommandDelete, query)

	if q == nil {
		t.Error("Query is nil")
	}

	if q.ID() != "1" {
		t.Errorf("Expected ID to be 1, got %s", q.ID())
	}

	if q.Type() != command.CommandDelete {
		t.Errorf("Expected type to be DELETE, got %s", q.Type())
	}

	if q.Key() != "key" {
		t.Errorf("Expected key to be key, got %s", q.Key())
	}

	if q.Value() != nil {
		t.Error("Value is not nil")
	}

	if q.TTL() != nil {
		t.Error("TTL is not nil")
	}
}
