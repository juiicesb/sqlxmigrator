//go:build postgresql
// +build postgresql

package sqlxmigrator

import (
	_ "github.com/lib/pq"
)

func init() {
	databases = append(databases, database{
		name:    "postgres",
		connEnv: "PG_CONN_STRING",
	})
}
