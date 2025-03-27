package hxsqlfiltering

import (
	"strconv"
	"strings"
)

func BuildQuery(base string, filters ...Operation) (string, []any) {
	if len(filters) == 0 {
		return base, []any{}
	}

	args := make([]any, len(filters), len(filters))
	clauses := make([]string, len(filters), len(filters))

	for ix, f := range filters {
		clauses[ix] = f.Operation() + strconv.Itoa(ix+1) + f.CloseOperation()
		args[ix] = f.GetArguments()
	}

	return base + _Where + strings.Join(clauses, _And), args
}
