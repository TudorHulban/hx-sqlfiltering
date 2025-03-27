package hxsqlfiltering

type FilterSmallerOrEqual struct {
	Arguments any

	Column string
}

func (f FilterSmallerOrEqual) Operation() string {
	return f.Column + " <= $"
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
