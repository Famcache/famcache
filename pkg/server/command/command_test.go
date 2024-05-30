package command

import "testing"

func TestNewCommand_GET(t *testing.T) {
	com, ok := NewCommand("id GET key")

	if ok != nil {
		t.Error("Expected no error")
	}
	if com == nil {
		t.Error("Expected a command")
	}

	if !com.IsStoreCommand() {
		t.Error("Expected a store command")
	}

	if com.cType != CommandGet {
		t.Error("Expected a GET command")
	}

	if com.ToStoreCommand().Key != "key" {
		t.Error("Expected key to be 'key'")
	}

	if com.ToStoreCommand().Value != nil {
		t.Error("Expected value to be nil")
	}

	if com.ToStoreCommand().ID != "id" {
		t.Error("Expected id to be 'id'")
	}
}

func TestNewCommand_SET(t *testing.T) {
	com, ok := NewCommand("id SET key value 0")

	if ok != nil {
		t.Error("Expected no error")
	}
	if com == nil {
		t.Error("Expected a command")
	}

	if !com.IsStoreCommand() {
		t.Error("Expected a store command")
	}

	if com.cType != CommandSet {
		t.Error("Expected a SET command")
	}

	if com.ToStoreCommand().Key != "key" {
		t.Error("Expected key to be 'key'")
	}

	if *com.ToStoreCommand().Value != "value" {
		t.Error("Expected value to be 'value'")
	}

	if *com.ToStoreCommand().TTL != 0 {
		t.Error("Expected ttl to be 0")
	}

	if com.ToStoreCommand().ID != "id" {
		t.Error("Expected id to be 'id'")
	}
}

func TestNewCommand_DELETE(t *testing.T) {
	com, ok := NewCommand("id DELETE key")

	if ok != nil {
		t.Error("Expected no error")
	}
	if com == nil {
		t.Error("Expected a command")
	}

	if !com.IsStoreCommand() {
		t.Error("Expected a store command")
	}

	if com.cType != CommandDelete {
		t.Error("Expected a DELETE command")
	}

	if com.ToStoreCommand().Key != "key" {
		t.Error("Expected key to be 'key'")
	}

	if com.ToStoreCommand().Value != nil {
		t.Error("Expected value to be nil")
	}

	if com.ToStoreCommand().ID != "id" {
		t.Error("Expected id to be 'id'")
	}
}

func TestNewCommand_Publish(t *testing.T) {
	com, ok := NewCommand("id PUBLISH channel message")

	if ok != nil {
		t.Error("Expected no error")
	}
	if com == nil {
		t.Error("Expected a command")
	}

	if com.cType != CommandPublish {
		t.Error("Expected a PUBLISH command")
	}

	if com.ToPubsubCommand().Topic != "channel" {
		t.Error("Expected channel to be 'channel'")
	}

	if com.ToPubsubCommand().Data != "message" {
		t.Error("Expected message to be 'message'")
	}

	if com.ToPubsubCommand().ID != "id" {
		t.Error("Expected id to be 'id'")
	}
}

func TestNewCommand_Subscribe(t *testing.T) {
	com, ok := NewCommand("id SUBSCRIBE channel")

	if ok != nil {
		t.Error("Expected no error")
	}
	if com == nil {
		t.Error("Expected a command")
	}

	if com.cType != CommandSubscribe {
		t.Error("Expected a SUBSCRIBE command")
	}

	if com.ToPubsubCommand().Topic != "channel" {
		t.Error("Expected channel to be 'channel'")
	}

	if com.ToPubsubCommand().Data != "" {
		t.Error("Expected message to be empty")
	}

	if com.ToPubsubCommand().ID != "id" {
		t.Error("Expected id to be 'id'")
	}
}

func TestNewCommand_Unsubscribe(t *testing.T) {
	com, error := NewCommand("id UNSUBSCRIBE channel")

	if error != nil {
		t.Error("Expected no error")
	}
	if com == nil {
		t.Error("Expected a command")
	}

	if com.cType != CommandUnsubscribe {
		t.Error("Expected a UNSUBSCRIBE command")
	}

	if com.ToPubsubCommand().Topic != "channel" {
		t.Error("Expected channel to be 'channel'")
	}

	if com.ToPubsubCommand().Data != "" {
		t.Error("Expected message to be empty")
	}

	if com.ToPubsubCommand().ID != "id" {
		t.Error("Expected id to be 'id'")
	}
}

func TestNewCommand_INVALID(t *testing.T) {
	com, ok := NewCommand("id INVALID key")

	if ok == nil {
		t.Error("Expected an error")
	}
	if com != nil {
		t.Error("Expected no command")
	}
}
