package hxsqlfiltering

type Operation interface {
	Operation(number int) string
	CloseOperation() string
	GetArguments() any
	IsOrdering() bool
	IsLimit() bool
}

var _ Operation = FilterEqual{}
var _ Operation = FilterExists{}
var _ Operation = FilterLike{}
var _ Operation = FilterSmallerOrEqual{}
var _ Operation = FilterGreaterOrEqual{}

var _ Operation = FilterOrderByColumn{}
var _ Operation = FilterOrderByColumns{}

var _ Operation = FilterPagination{}
