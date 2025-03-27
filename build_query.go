package hxsqlfiltering

import (
	"strconv"
	"strings"
)

func BuildQuery(base string, filters ...Operation) (string, []any) {
	if len(filters) == 0 {
		return base, []any{}
	}

	args := make([]any, 0)
	clauses := make([]string, 0)

	for ix, f := range filters {
		if f.IsOrdering() {
			return base +
					ternary(len(args) > 0, _Where, "") +
					strings.Join(clauses, _And) +
					f.CloseOperation(),
				args
		}

		operation := f.Operation()

		clause := f.Operation() + strconv.Itoa(ix+1) + ternary(
			strings.Contains(operation, "("),

			")",
			"",
		)

		clauses = append(clauses, clause)

		args = append(args, f.GetArguments())
	}

	return base + _Where + strings.Join(clauses, _And), args
}
