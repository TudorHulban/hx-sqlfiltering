# hx-sqlfiltering

Simple dynamic SQL filtering for Go.

```go
var args []any

f := Filter{
    Args: &args,
    Column: "name",
    Value: "john%",
    Valid: true,
}

sql, args := BuildQuery("SELECT * FROM users", []Filter{f})
// "SELECT * FROM users WHERE name ILIKE $1", ["john%"]
```
