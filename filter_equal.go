package hxsqlfiltering

type FilterEqual struct {
	Arguments any

	Column string
}

func (f FilterEqual) Operation() string {
	return f.Column + " = $"
}

func (f FilterEqual) CloseOperation() string {
	return ""
}

func (f FilterEqual) GetArguments() any {
	return f.Arguments
}
