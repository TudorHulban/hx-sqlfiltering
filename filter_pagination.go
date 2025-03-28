package hxsqlfiltering

import (
	"database/sql"
	"strconv"
)

type FilterPagination struct {
	First uint8
	After sql.NullInt16
}

func (f FilterPagination) Operation(number int) string {
	if f.First == 0 {
		return ""
	}

	return " limit " +
		strconv.FormatUint(uint64(f.First), 10) +
		ternary(
			f.After.Valid,

			" offset "+strconv.FormatInt(int64(f.After.Int16), 10),
			"",
		) +
		";"
}

func (f FilterPagination) CloseOperation() string {
	return ""
}

func (f FilterPagination) GetArguments() any {
	return nil
}

func (f FilterPagination) IsOrdering() bool {
	return false
}

func (f FilterPagination) IsLimit() bool {
	return true
}
