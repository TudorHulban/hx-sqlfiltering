package hxsqlfiltering

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSQLFilteringWithSQLite(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY,
			name TEXT,
			email TEXT,
			age INTEGER,
			is_active BOOLEAN
		)
	`)
	require.NoError(t, err)

	_, err = db.Exec(`
		INSERT INTO users (id, name, email, age, is_active) VALUES
		(123, 'John Doe', 'john@example.com', 30, 1),
		(124, 'Jane Smith', 'jane@example.com', 25, 1),
		(125, 'Bob Johnson', 'bob@example.com', 40, 0)
	`)
	require.NoError(t, err)

	tests := []struct {
		name     string
		base     string
		filters  []Operation
		wantSql  string
		wantArgs []any
		wantRows int
	}{
		{
			name: "single equality filter",
			base: "select * from users",
			filters: []Operation{
				FilterEqual{Column: "id", Arguments: int64(123)},
			},
			wantSql:  "select * from users where id = $1",
			wantArgs: []any{int64(123)},
			wantRows: 1,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				generatedSQL, generatedArgs := BuildQuery(
					tt.base,
					tt.filters...,
				)
				require.Equal(t, tt.wantSql, generatedSQL)
				require.Equal(t, tt.wantArgs, generatedArgs)

				rows, err := db.Query(generatedSQL, generatedArgs...)
				require.NoError(t, err)
				defer rows.Close()

				var rowsCount int

				for rows.Next() {
					rowsCount++
				}
				require.NoError(t, rows.Err())
				require.Equal(t, tt.wantRows, rowsCount)
			})
	}
}
