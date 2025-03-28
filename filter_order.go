package hxsqlfiltering

import (
	"strings"
)

type FilterOrderByColumn struct {
	Column     string
	Descending bool
}

func (f FilterOrderByColumn) Operation(_ int) string {
	if f.Descending {
		return sprintf("order by %s desc", f.Column)
	}

	return sprintf("order by %s asc", f.Column)
}

func (f FilterOrderByColumn) CloseOperation() string {
	return ""
}

func (f FilterOrderByColumn) GetArguments() any {
	return nil
}

func (f FilterOrderByColumn) IsOrdering() bool {
	return true
}

func (f FilterOrderByColumn) IsLimit() bool {
	return false
}

type FilterOrderByColumns struct {
	Columns    []string
	Descending []bool
}

func (f FilterOrderByColumns) Operation(number int) string {
	return ""
}

func (f FilterOrderByColumns) CloseOperation() string {
	var clauses []string

	for i, col := range f.Columns {
		dir := "asc"

		if i < len(f.Descending) && f.Descending[i] {
			dir = "desc"
		}

		clauses = append(clauses, sprintf("%s %s", col, dir))
	}

	return " order by " + strings.Join(clauses, ", ")
}

func (f FilterOrderByColumns) GetArguments() any {
	return nil
}

func (f FilterOrderByColumns) IsOrdering() bool {
	return true
}

func (f FilterOrderByColumns) IsLimit() bool {
	return false
}
