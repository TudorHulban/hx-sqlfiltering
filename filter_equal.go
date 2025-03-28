package hxsqlfiltering

import "strconv"

type FilterEqual struct {
	Arguments any

	Column string
}

func (f FilterEqual) Operation(number int) string {
	return f.Column + " = $" + strconv.Itoa(number+1)
}

func (f FilterEqual) CloseOperation() string {
	return ""
}

func (f FilterEqual) GetArguments() any {
	return f.Arguments
}

func (f FilterEqual) IsOrdering() bool {
	return false
}

func (f FilterEqual) IsLimit() bool {
	return false
}
