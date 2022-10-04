package surrealdb_go_utils

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/goccy/go-json"
)

var ErrInvalidThing = fmt.Errorf("invalid thing")

// NewThing Creates a new thing
// If we pass a single value, it will split it and validate.
// For example:
//
//	// Single thing value string
//	thing := NewThing("table:id")
//	// Table and id separately
//	thing := NewThing("table", "id")
func NewThing(values ...string) *Thing {
	length := len(values)
	if length == 0 || length > 2 {
		panic(ErrInvalidThing)
	}

	thing := &Thing{}

	if length == 1 {
		thing.value = values[0]
	}

	if length == 2 {
		thing.table = values[0]
		thing.id = values[1]
		thing.value = thing.table + ":" + thing.id
	}

	thing.validate()

	return thing
}

type Thing struct {
	table string
	id    string
	value string
}

func (thing *Thing) validate() {
	str, _ := url.QueryUnescape(thing.value)
	splitterIndex := strings.Index(str, ":")
	if splitterIndex == -1 {
		panic(ErrInvalidThing)
	}

	thing.table = str[:splitterIndex]
	thing.id = str[splitterIndex+1:]
	thing.value = str
}

// Table Get the table value
func (thing *Thing) Table() string {
	return thing.table
}

// Id Get the id value
func (thing *Thing) Id() string {
	return thing.id
}

// Value Get the full value, for example `table:id` or `user:one`
func (thing *Thing) Value() string {
	return thing.value
}

// AsTypeThing Get the value for the database, for example `type::thing('table','id')`
func (thing *Thing) AsTypeThing() string {
	return fmt.Sprintf("type::thing('%s','%s')", thing.table, thing.id)
}

// AsTypeThingParam Get the value for the database, like AsTypeThing, but it will use a param for the value instead.
// This will automatically add the $ to the param, for example:
//
//	thing := NewThing("user", "31367126")
//	thing.AsTypeThingParam("id") // type::thing('user', $id)
//	db.Query(fmt.Sprintf("select * from %s", thing.AsTypeThingParam("id")), map[string]any{
//	    "id": thing.Id(),
//	})
//
// Will output the query:
//
//	select * from type::thing('user', $id)
func (thing *Thing) AsTypeThingParam(paramName string) string {
	return fmt.Sprintf("type::thing('%s', %s)", thing.table, "$"+paramName)
}

func (thing *Thing) UnmarshalJSON(data []byte) error {
	_, dataType, _, err := jsonparser.Get(data)
	if err != nil {
		return err
	}

	if dataType != jsonparser.String {
		return ErrInvalidThing
	}

	value, err := jsonparser.GetString(data)
	if err != nil {
		return err
	}

	*thing = *NewThing(value)

	return nil
}

func (thing *Thing) MarshalJSON() ([]byte, error) {
	return json.Marshal(thing.value)
}
