package hxsqlfiltering

type FilterGreaterOrEqual struct {
	Arguments any

	Column string
}

func (f FilterGreaterOrEqual) Operation() string {
	return f.Column + " >= $"
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
