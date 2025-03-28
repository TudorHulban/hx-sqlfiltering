package hxsqlfiltering

import (
	"strconv"
	"strings"
)

type FilterExists struct {
	Arguments any

	ColumnJoin string
	TableJoin  string
	SubColumn  string
}

func (f FilterExists) Operation(number int) string {
	var b strings.Builder

	b.WriteString("exists (select 1 from ")
	b.WriteString(f.TableJoin)
	b.WriteString(" jt where jt.")
	b.WriteString(f.SubColumn)
	b.WriteString(" = ")
	b.WriteString(f.ColumnJoin)
	b.WriteString(" and jt.")
	b.WriteString(f.SubColumn)
	b.WriteString(" = $")
	b.WriteString(strconv.Itoa(number + 1))
	b.WriteString(")")

	return b.String()
}

func (f FilterExists) CloseOperation() string {
	return ")"
}

func (f FilterExists) GetArguments() any {
	return f.Arguments
}

func (f FilterExists) IsOrdering() bool {
	return false
}

func (f FilterExists) IsLimit() bool {
	return false
}
