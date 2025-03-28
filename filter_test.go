package hxsqlfiltering

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildQuery(t *testing.T) {
	tests := []struct {
		name     string
		base     string
		filters  []Operation
		wantSql  string
		wantArgs []any
	}{
		{
			name:     "1. no filters",
			base:     "select * from users",
			filters:  []Operation{},
			wantSql:  "select * from users",
			wantArgs: []any{},
		},
		{
			name: "2. single equality filter",
			base: "select * from users",
			filters: []Operation{
				FilterEqual{Column: "id", Arguments: int64(123)},
			},
			wantSql:  "select * from users where id = $1",
			wantArgs: []any{int64(123)},
		},
		{
			name: "3. like filter",
			base: "select * from users",
			filters: []Operation{
				FilterLike{Column: "name", Arguments: "john%"},
			},
			wantSql:  "select * from users where name ilike $1",
			wantArgs: []any{"john%"},
		},
		{
			name: "4. subquery filter",
			base: "select * from tickets",
			filters: []Operation{
				FilterExists{ColumnJoin: "id", TableJoin: "blocks", SubColumn: "id", Arguments: int64(456)},
			},
			wantSql:  "select * from tickets where exists (select 1 from blocks jt where jt.id = id and jt.id = $1)",
			wantArgs: []any{int64(456)},
		},
		{
			name: "5. multiple filters",
			base: "select * from tickets",
			filters: []Operation{
				FilterEqual{Column: "id", Arguments: int64(123)},
				FilterLike{Column: "name", Arguments: "issue%"},
				FilterExists{ColumnJoin: "ticket_id", TableJoin: "events", SubColumn: "id", Arguments: int64(789)},
			},
			wantSql:  "select * from tickets where id = $1 and name ilike $2 and exists (select 1 from events jt where jt.id = ticket_id and jt.id = $3)",
			wantArgs: []any{int64(123), "issue%", int64(789)},
		},
		{
			name: "6. greater than filter",
			base: "select * from users",
			filters: []Operation{
				FilterSmallerOrEqual{Column: "age", Arguments: int64(30)},
			},
			wantSql:  "select * from users where age <= $1",
			wantArgs: []any{int64(30)},
		},
		{
			name: "7. greater or equal filter",
			base: "select * from users",
			filters: []Operation{
				FilterGreaterOrEqual{Column: "age", Arguments: int64(30)},
			},
			wantSql:  "select * from users where age >= $1",
			wantArgs: []any{int64(30)},
		},
		{
			name: "8. multi-column ordering with no additional filtering",
			base: "select * from tickets",
			filters: []Operation{
				FilterOrderByColumns{
					Columns:    []string{"priority", "created_at"},
					Descending: []bool{false, true},
				},
			},
			wantSql:  "select * from tickets order by priority asc, created_at desc",
			wantArgs: []any{},
		},
		{
			name: "9. multi-column ordering with mixed directions",
			base: "select * from tickets",
			filters: []Operation{
				FilterExists{
					ColumnJoin: "project_id",
					TableJoin:  "projects",
					SubColumn:  "id",
					Arguments:  100,
				},
				FilterOrderByColumns{
					Columns:    []string{"priority", "created_at"},
					Descending: []bool{false, true},
				},
			},
			wantSql:  "select * from tickets where exists (select 1 from projects jt where jt.id = project_id and jt.id = $1) order by priority asc, created_at desc",
			wantArgs: []any{100},
		},
		{
			name: "10. pagination with limit only",
			base: "select * from users",
			filters: []Operation{
				FilterPagination{
					First: 10,
					After: sql.NullInt16{Valid: false},
				},
			},
			wantSql:  "select * from users limit 10;",
			wantArgs: []any{},
		},
		{
			name: "11. pagination with limit and offset",
			base: "select * from users",
			filters: []Operation{
				FilterPagination{
					First: 5,
					After: sql.NullInt16{Int16: 10, Valid: true},
				},
			},
			wantSql:  "select * from users limit 5 offset 10;",
			wantArgs: []any{},
		},
		{
			name: "12. pagination with limit and offset",
			base: "select * from users",
			filters: []Operation{
				FilterEqual{Column: "id", Arguments: int64(123)},
				FilterPagination{
					First: 5,
					After: sql.NullInt16{Int16: 10, Valid: true},
				},
			},
			wantSql:  "select * from users where id = $1 limit 5 offset 10;",
			wantArgs: []any{int64(123)},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				gotSql, gotArgs := BuildQuery(tt.base, tt.filters...)
				require.Equal(t,
					tt.wantSql,
					gotSql,
				)

				if !reflect.DeepEqual(gotArgs, tt.wantArgs) {
					t.Errorf("BuildQuery() args = %v, want %v", gotArgs, tt.wantArgs)
				}
			},
		)
	}
}
