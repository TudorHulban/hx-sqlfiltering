package hxsqlfiltering

import "strconv"

type FilterLike struct {
	Arguments string

	Column string
}

func (f FilterLike) Operation(number int) string {
	return f.Column + " ilike $" + strconv.Itoa(number+1)
}

func (f FilterLike) CloseOperation() string {
	return ""
}

func (f FilterLike) GetArguments() any {
	return f.Arguments
}

func (f FilterLike) IsOrdering() bool {
	return false
}
