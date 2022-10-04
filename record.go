package surrealdb_go_utils

import (
	"github.com/buger/jsonparser"
	"github.com/goccy/go-json"
)

type RecordType string

var (
	RecordTypeId     RecordType = "id"
	RecordTypeRecord RecordType = "record"
)

type Record[T any] struct {
	id    string
	value T
	isSet bool
	kind  RecordType
}

func (record *Record[T]) UnmarshalJSON(data []byte) error {
	_, dataType, _, err := jsonparser.Get(data)
	if err != nil {
		return err
	}

	dat := *new(Record[T])

	if dataType == jsonparser.String {
		dat.kind = RecordTypeId
		dat.id, err = jsonparser.GetString(data)
		if err != nil {
			return err
		}
		dat.isSet = true

		*record = dat

		return nil
	}
	if dataType == jsonparser.Object {
		var item T = *new(T)
		err = json.Unmarshal(data, &item)
		if err != nil {
			return err
		}
		dat.value = item
		dat.kind = RecordTypeRecord
		dat.isSet = true

		dat.id, err = jsonparser.GetString(data, "id")
		if err != nil {
			return err
		}

		*record = dat

		return nil
	}

	return nil
}

func (record *Record[T]) MarshalJSON() ([]byte, error) {
	if record.IsId() {
		return []byte(`"` + record.id + `"`), nil
	}

	return json.Marshal(record.value)
}

// IsId returns true if the record is an id
func (record *Record[T]) IsId() bool {
	return record.kind == RecordTypeId
}

// Id returns the id of the record(used when we haven't fetched the record yet)
func (record *Record[T]) Id() string {
	return record.id
}

// IsRecord returns true if we're using record rather than id
func (record *Record[T]) IsRecord() bool {
	return record.kind == RecordTypeRecord
}

// Record returns the record
func (record *Record[T]) Record() T {
	return record.value
}
