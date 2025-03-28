package hxsqlfiltering

import "strconv"

type FilterSmallerOrEqual struct {
	Arguments any

	Column string
	Strict bool
}

func (f FilterSmallerOrEqual) Operation(number int) string {
	if f.Strict {
		return f.Column + " = $" + strconv.Itoa(number+1)
	}

	return f.Column + " <= $" + strconv.Itoa(number+1)
}

func (f FilterSmallerOrEqual) CloseOperation() string {
	return ""
}

func (f FilterSmallerOrEqual) GetArguments() any {
	return f.Arguments
}

func (f FilterSmallerOrEqual) IsOrdering() bool {
	return false
}
