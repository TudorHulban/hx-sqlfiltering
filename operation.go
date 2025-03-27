package hxsqlfiltering

type Operation interface {
	Operation() string
	CloseOperation() string
	GetArguments() any
}
