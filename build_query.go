package hxsqlfiltering

import (
	"strings"
)

func BuildQuery(base string, filters ...Operation) (string, []any) {
	if len(filters) == 0 {
		return base, []any{}
	}

	args := make([]any, 0)
	clauses := make([]string, 0)
	query := base

	for ix, f := range filters {
		if f.IsOrdering() {
			query = base +
				ternary(len(args) > 0, _Where, "") +
				strings.Join(clauses, _And) +
				f.CloseOperation()

			if ix == len(filters)-1 {
				return query,
					args
			}

			continue
		}

		if f.IsLimit() {
			// Include clauses before appending limit
			if len(clauses) > 0 {
				query = base + _Where + strings.Join(clauses, _And)
			}

			query = query + f.Operation(ix)

			return query,
				args
		}

		// Only non-limit filters go into clauses
		clauses = append(
			clauses,
			f.Operation(ix),
		)

		if arguments := f.GetArguments(); arguments != nil {
			args = append(args, arguments)
		}
	}

	return base + _Where + strings.Join(clauses, _And),
		args
}
