package hxsqlfiltering

type FilterLike struct {
	Arguments string

	Column string
}

func (f FilterLike) Operation() string {
	return f.Column + " ilike $"
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
