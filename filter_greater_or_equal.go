package hxsqlfiltering

import "strconv"

type FilterGreaterOrEqual struct {
	Arguments any

	Column string
	Strict bool
}

func (f FilterGreaterOrEqual) Operation(number int) string {
	if f.Strict {
		return f.Column + " > $" + strconv.Itoa(number+1)
	}

	return f.Column + " >= $" + strconv.Itoa(number+1)
}

func (f FilterGreaterOrEqual) CloseOperation() string {
	return ""
}

func (f FilterGreaterOrEqual) GetArguments() any {
	return f.Arguments
}

func (f FilterGreaterOrEqual) IsOrdering() bool {
	return false
}

func (f FilterGreaterOrEqual) IsLimit() bool {
	return false
}
