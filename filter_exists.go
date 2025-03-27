package hxsqlfiltering

import "strings"

type FilterExists struct {
	Arguments any

	ColumnJoin string
	Table      string
	SubColumn  string
}

func (f FilterExists) Operation() string {
	var b strings.Builder

	b.WriteString("exists (select 1 from ")
	b.WriteString(f.Table)
	b.WriteString(" b where b.")
	b.WriteString(f.SubColumn)
	b.WriteString(" = ")
	b.WriteString(f.ColumnJoin)
	b.WriteString(" and b.")
	b.WriteString(f.SubColumn)
	b.WriteString(" = $")

	return b.String()
}

func (f FilterExists) CloseOperation() string {
	return ")"
}

func (f FilterExists) GetArguments() any {
	return f.Arguments
}
