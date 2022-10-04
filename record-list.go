package surrealdb_go_utils

import (
	"github.com/buger/jsonparser"
	"github.com/goccy/go-json"
	"golang.org/x/exp/maps"
)

type ListType string

var (
	ListTypeIds     ListType = "ids"
	ListTypeObjects ListType = "objects"
)

type RecordList[T any | string] struct {
	items    []T
	ids      map[string]*T
	listType ListType
}

func (d *RecordList[T]) MarshalJSON() ([]byte, error) {
	if d.IsResolved() {
		return json.Marshal(d.items)
	}

	return json.Marshal(d.Ids())
}

func (d *RecordList[T]) UnmarshalJSON(data []byte) error {
	_, dataType, _, err := jsonparser.Get(data)
	if err != nil {
		return err
	}

	if dataType == jsonparser.Array {
		dat := *new(RecordList[T])
		dat.ids = make(map[string]*T)
		dat.items = make([]T, 0)
		dat.listType = ""

		jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			if dataType == jsonparser.String {
				val, err := jsonparser.ParseString(value)
				if err != nil {
					panic(err)
				}
				dat.ids[val] = nil
				dat.listType = ListTypeIds
			} else {
				var item T = *new(T)
				err = json.Unmarshal(value, &item)
				if err != nil {
					return
				}
				dat.items = append(dat.items, item)
				dat.listType = ListTypeObjects
			}
		})

		*d = dat

		return nil
	}

	return nil
}

// IsResolved check if the list has been resolved/initiated
func (d *RecordList[T]) IsResolved() bool {
	return d.listType != ListTypeIds
}

// Items returns the items in the list
func (d *RecordList[T]) Items() []T {
	return d.items
}

// IdMap returns the ids in the list(this would be used when we haven't fetched the records)
func (d *RecordList[T]) IdMap() map[string]*T {
	return d.ids
}

// Ids return an array of all ids in the list
func (d *RecordList[T]) Ids() []string {
	return maps.Keys(d.ids)
}
