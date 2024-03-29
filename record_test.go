package surrealdb_go_utils

import (
	"testing"

	"github.com/goccy/go-json"
)

func TestFetchedRecordUnmarshal(t *testing.T) {
	type subrec struct {
		Id *Thing `json:"id"`
	}
	type rec struct {
		Id    *Thing `json:"id"`
		Value *Record[subrec]
	}

	jsonWithFetch := `{"id": "user:one", "value": {"id": "book:one"}}`
	var r rec
	err := json.Unmarshal([]byte(jsonWithFetch), &r)
	if err != nil {
		t.Error(err)
	}

	if r.Value.IsRecord() == false {
		t.Error("Invalid value")
	}

	if r.Value.Id() != "book:one" {
		t.Error("Invalid value")
	}

	if r.Value.Record().Id.Value() != "book:one" {
		t.Error("Invalid value")
	}
}

func TestRecordPointerUnmarshal(t *testing.T) {
	type subrec struct {
		Id *Thing `json:"id"`
	}
	type rec struct {
		Id    *Thing `json:"id"`
		Value *Record[subrec]
	}

	jsonWithFetch := `{"id": "user:one", "value": "book:one"}`
	var r rec
	err := json.Unmarshal([]byte(jsonWithFetch), &r)
	if err != nil {
		t.Error(err)
	}

	if r.Value.IsId() == false {
		t.Error("Invalid value")
	}

	if r.Value.Id() != "book:one" {
		t.Error("Invalid value")
	}
}

func TestRecordCreation(t *testing.T) {
	type subrec struct {
		Id *Thing `json:"id"`
	}
	type rec struct {
		Id    *Thing `json:"id"`
		Value *Record[subrec]
	}

	r := rec{
		Id:    NewThing("user:one"),
		Value: NewRecord[subrec](NewThing("book:one")),
	}

	if r.Value.IsId() == false {
		t.Error("Invalid value")
	}

	if r.Value.Id() != "book:one" {
		t.Error("Invalid value")
	}

	r = rec{
		Id:    NewThing("user:one"),
		Value: NewRecord[subrec](subrec{Id: NewThing("book:one")}),
	}

	if r.Value.IsRecord() == false {
		t.Error("Invalid value")
	}

	if r.Value.Id() != "book:one" {
		t.Error("Invalid value")
	}
}
func TestCreatingRecordFromThing(t *testing.T) {
	type subrec struct {
		Id *Thing `json:"id"`
	}
	type rec struct {
		Id    *Thing `json:"id"`
		Value *Record[subrec]
	}

	r := rec{
		Id:    NewThing("user:one"),
		Value: NewRecord[subrec](NewThing("book:one")),
	}

	newRecord := NewRecord[subrec](r.Id)

	if newRecord.IsId() == false {
		t.Error("Invalid value")
	}

	if newRecord.Id() != "user:one" {
		t.Error("Invalid value")
	}

}
