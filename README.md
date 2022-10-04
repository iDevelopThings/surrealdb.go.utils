# SurrealDB Go Utils

A basic little package with some useful structs for working with surreal

### Thing:

```go
package main

import (
	"github.com/idevelopthings/surrealdb.go.utils"
)

type User struct {
	ID surrealdb_go_utils.Thing `json:"id"`
}
```

Thing will allow you to the regular surrealdb id formatting, and take out the Table or Value with ease.
We can also use this for situations in queries.

It will also unescape the value for you, so when it's passed via an api endpoint as a normal id, like "user:u4y127h7fh78h1", it will automatically be unescaped.

### Record:

This is useful for when you have a type for your Table, which will have a pointer value, for example. A Book with an Author:

```go 
package main

import (
	"github.com/idevelopthings/surrealdb_go_utils"
)

type Author struct {
    ID      surrealdb_go_utils.Thing `json:"id"`
}
type Book struct {
    ID      surrealdb_go_utils.Thing `json:"id"`
    Author  surrealdb_go_utils.Record[Author] `json:"author"`
}

```

When you pull the Book from the db as normal, it will only return an id string, something like `{"id": "book:u4y127h7fh78h1", "author": "author:jsdfhf8f2h7fh"}`.

When you also fetch "author", this usually will cause issues unmarsalling the json.

But this "Record" will handle both situations, and marshal to json as the correct type

```go
var book *Book

book.Author.IsId()     // Returns true if the value is an id(record pointer)
book.Author.Id()       // Returns the ID of the record 
book.Author.IsRecord() // Returns true if it's the Author value/record
book.Author.Record()   // Returns the Author record(if the query included `fetch author` for example)

```

But when you use the Record, it will automatically unescape the id, and then pull the Author from the db, and return the Author struct.

### RecordList

This works the same as "Record" above, but for an array instead.