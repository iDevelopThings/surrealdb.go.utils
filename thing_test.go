package surrealdb_go_utils

import (
	"testing"

	"github.com/goccy/go-json"
)

func TestThing(t *testing.T) {
	thing := NewThing("user:one")

	if thing.Table() != "user" {
		t.Error("Invalid table")
	}

	if thing.Id() != "one" {
		t.Error("Invalid id")
	}

	if thing.Value() != "user:one" {
		t.Error("Invalid value")
	}

	thing = NewThing("user", "one")

	if thing.Table() != "user" {
		t.Error("Invalid table")
	}

	if thing.Id() != "one" {
		t.Error("Invalid id")
	}

	if thing.Value() != "user:one" {
		t.Error("Invalid value")
	}
}

func TestThingJsonDecoding(t *testing.T) {
	type rec struct {
		Id *Thing `json:"id"`
	}
	jsonValue := `{"id":"user:one"}`

	var r rec
	err := json.Unmarshal([]byte(jsonValue), &r)
	if err != nil {
		t.Error(err)
	}

	if r.Id.Table() != "user" {
		t.Error("Invalid table")
	}

	if r.Id.Id() != "one" {
		t.Error("Invalid id")
	}

	if r.Id.Value() != "user:one" {
		t.Error("Invalid value")
	}

}

func TestThingJsonEncoding(t *testing.T) {
	type rec struct {
		Id *Thing `json:"id"`
	}
	value := rec{Id: NewThing("user:one")}

	jsonValue, err := json.Marshal(value)
	if err != nil {
		t.Error(err)
	}

	if string(jsonValue) != `{"id":"user:one"}` {
		t.Error("Invalid json")
	}
}
