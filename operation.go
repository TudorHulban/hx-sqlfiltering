package hxsqlfiltering

type Operation interface {
	Operation() string
	CloseOperation() string
	GetArguments() any
	IsOrdering() bool
}

var _ Operation = FilterEqual{}
var _ Operation = FilterExists{}
var _ Operation = FilterLike{}
var _ Operation = FilterSmallerOrEqual{}
var _ Operation = FilterGreaterOrEqual{}

var _ Operation = FilterOrderByColumn{}
var _ Operation = FilterOrderByColumns{}
